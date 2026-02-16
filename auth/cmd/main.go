package main

import (
	"clirzy/auth/internal/infrastructure/builder"
	"clirzy/auth/internal/infrastructure/config"
	pkgapp "clirzy/pkg/app"
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
