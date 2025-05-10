package main

import (
	"bankapp2/bootstrapper"
	"context"
	"log"

	"github.com/robfig/cron"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := cron.New()
	defer c.Stop()

	app := bootstrapper.New()
	if err := app.RunAPI(ctx); err != nil {
		log.Fatal("failed to start app")
	}
	if err := app.RunCron(ctx, c); err != nil {
		log.Fatal("failed to start cron")
	}

	defer app.StopKafka()

}
