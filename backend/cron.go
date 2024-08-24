package main

import (
	"binalyze-test/setup"
	"context"

	"github.com/robfig/cron/v3"
)

func RunSchedule(services *setup.ServiceDependencies) {
	ctx := context.Background()
	s := cron.New()

	s.AddFunc("@every 30s", func() {
		services.Logger.Info("Running cron service")
		services.ProcessService.FetchAndInsertProcess(ctx)
	})

	s.Start()
}
