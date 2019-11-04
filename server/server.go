package server

import (
	"net/http"
	"reflect"
	"runtime"

	"github.com/8bitdogs/ruffe"
	"github.com/antonmashko/log"
)

type Server struct {
	logger     log.Logger
	rs         *ruffe.Server
	handler    *ruffe.Middleware
	useRecover bool
}

func New() *Server {
	return &Server{
		rs:     ruffe.New(),
		logger: log.DefaultLogger.Copy("server"),
	}
}

func (s *Server) Handle(pattern, method string, h ruffe.Handler) {
	s.logger.Infof("adding pattern=%s method=%s handler=%s",
		pattern, method, runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
	if s.useRecover {
		h = s.recoverMiddleware(h)
	}
	s.rs.Handle(pattern, method, h)
}

func (s *Server) HandleFunc(pattern, method string, h func(ruffe.Context) error) {
	s.Handle(pattern, method, ruffe.HandlerFunc(h))
}

func (s *Server) UseAccessLog() log.Logger {
	al := log.DefaultLogger.Copy("request")
	s.rs.UseFunc(func(ctx ruffe.Context) error {
		r := ctx.Request()
		al.Infoln(r.Method, r.RemoteAddr, r.RequestURI, r.UserAgent(), r.Proto)
		return nil
	})
	return al
}

// TODO: implement me
// func (s *Server) UseResponseLog() log.Logger {
// 	return nil
// }

func (s *Server) UseRecover() {
	s.useRecover = true
}

func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, s.rs)
}

func (s *Server) recoverMiddleware(h ruffe.Handler) ruffe.Handler {
	return ruffe.HandlerFunc(func(ctx ruffe.Context) error {
		defer func() {
			if r := recover(); r != nil {
				s.logger.Errorf("panic recovered. cause=%v", r)
				ctx.Result(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				return
			}
		}()
		return h.Handle(ctx)
	})
}
