package main

import (
	pkgapp "clirzy/pkg/app"
	"clirzy/user/internal/infrastructure/builder"
	"clirzy/user/internal/infrastructure/config"
	"context"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app, err := builder.BuildApp(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := pkgapp.Run(context.Background(), app); err != nil {
		log.Fatal(err)
	}
}
