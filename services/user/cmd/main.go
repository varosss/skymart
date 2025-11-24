package main

import (
	"clirzy/pkg/config"
	"clirzy/services/user/internal/app"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("couldn't load app config")

		return
	}

	app, err := app.New(cfg)
	if err != nil {
		log.Fatalf("couldn't start app: %s", err.Error())

		return
	}

	app.Run()
}
