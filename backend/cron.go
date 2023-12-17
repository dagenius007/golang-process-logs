package main

import (
	"fmt"

	"binalyze-test/handlers"

	"github.com/robfig/cron/v3"
)

func RunSchedule() {
	fmt.Println("Fetching Process")
	s := cron.New()

	s.AddFunc("@every 10s", func() {
		fmt.Println("Fetching Process")
		handlers.FetchAndInsertProcess()
	})

	s.Start()
}
