package main

import (
	"os"
	"os/signal"
	"syscall"

	"webservice-template/database"
	"webservice-template/server"

	"github.com/8bitdogs/ruffe"
	"github.com/antonmashko/log"
	_ "github.com/lib/pq"
)

func main() {
	log.Info("starting application...")
	cfg, err := parse()
	if err != nil {
		log.Fatal(err)
	}

	log.Info(cfg.Database.ConnectionString())
	_, err = database.Connect(cfg.DatabaseDriveName, cfg.Database.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New()
	_ = srv.UseAccessLog()
	srv.HandleFunc("/test", "GET", func(ctx ruffe.Context) error {
		ctx.Result(200, "fine")
		return nil
	})

	go srv.ListenAndServe(cfg.Server.Addr)
	graceful()
}

func graceful() {
	// handling signals. blocking main goroutine
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-sigc
		log.Infoln("received signal ", s)
		switch s {
		case syscall.SIGTERM:
			// TODO: implement graceful shutdown
			return
		case syscall.SIGQUIT:
			return
		default:
			log.Infoln("ignoring signal ", s)
		}
	}
}
