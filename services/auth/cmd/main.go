package main

import (
	"log"

	userclient "clirzy/pkg/client/user"
	"clirzy/services/auth/internal/app"
	"clirzy/services/auth/internal/service"
)

func main() {
	userClient, err := userclient.New("user:50051")
	if err != nil {
		log.Fatalf("got fatal error: %s", err.Error())

		return
	}

	authService := service.NewAuthService(userClient)
	app := app.New(authService)

	app.Run()
}
