package main

import (
	"context"
	"flag"

	"github.com/chubin518/kestrel-layout-basic/pkg/config"
	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
	"github.com/chubin518/kestrel-layout-basic/pkg/logging"
)

func main() {
	flag.Parse()
	defer func() {
		if err := recover(); err != nil {
			logging.Errorf("Recovered err: %v", err)
		}
	}()
	// init confg
	err := config.Init()
	if err != nil {
		panic(err)
	}

	// init logging
	logging.InitWithOptions(config.Logging())

	// init config
	app, err := newApp()
	if err != nil {
		logging.Errorf("init app error: %v", err)
	}

	// graceful shutdown
	graceful.NewBuilderWithOptions(config.Server()).UseHooks(app.Hooks...).UseRoutes(app.Routes...).Build().RunWithContext(context.Background())
}
