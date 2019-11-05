package main

import (
	"os"
	"os/signal"
	"syscall"

	"webservice-template/database"
	"webservice-template/server"

	"github.com/antonmashko/log"
	_ "github.com/lib/pq"
)

func main() {
	log.Info("starting application...")
	cfg, err := parse()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Connect(cfg.DatabaseDriveName, cfg.Database.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	srv := server.New()
	srv.UseAccessLog()
	srv.RecoverPanic()

	// TODO: add handlers

	log.Infoln("starting server on", cfg.Server.Addr)
	go srv.ListenAndServe(cfg.Server.Addr)
	graceful()
}

func graceful() {
	// handling signals. blocking main gorqoutine
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
