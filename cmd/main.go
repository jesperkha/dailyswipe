package main

import (
	"log"
	"os"
	"syscall"

	"github.com/jesperkha/dailyswipe/config"
	"github.com/jesperkha/dailyswipe/server"
	"github.com/jesperkha/notifier"
)

func main() {
	notif := notifier.New()

	config := config.Load()
	s := server.New(config)
	s.Use(server.Logger)

	go s.ListenAndServe(notif)

	notif.NotifyOnSignal(os.Interrupt, syscall.SIGTERM)
	log.Println("shutdown")
}
