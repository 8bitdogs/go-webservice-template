package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/antonmashko/log"
)

func main() {
	log.Info("starting application...")
	_, err := parse()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: define and start all required services
	// ...

	// handling signals. blocking main goroutine
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-sigc
		log.Info("received signal ", s)
		switch s {
		case syscall.SIGTERM:
			// TODO: implement graceful shutdown
			return
		case syscall.SIGQUIT:
			return
		default:
			log.Info("ignoring signal ", s)
		}
	}
}
