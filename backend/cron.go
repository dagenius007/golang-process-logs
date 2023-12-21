package main

import (
	"log"

	"binalyze-test/handlers"

	"github.com/robfig/cron/v3"
)

func RunSchedule() {
	s := cron.New()

	s.AddFunc("@every 30s", func() {
		log.Println("Cron service running")
		handlers.FetchAndInsertProcess()
	})

	s.Start()
}
