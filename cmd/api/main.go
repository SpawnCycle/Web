package main

import (
	"context"
	"log"
	"smaash-web/internal/controllers"
	"smaash-web/internal/repository"
	"smaash-web/internal/server"
	"smaash-web/internal/services"
)

// TODO: refactor controllers so that the server depends on them, solvind coupling issue

func main() {
	appContext := context.Background()
	userStatsService := services.NewUserStatsService(
		repository.NewGormUserRepo(),
	)

	srv := server.NewServer(
		*controllers.NewUserController(userStatsService),
	).MountRoutes()

	if err := srv.Run(appContext); err != nil {
		log.Fatalf("There was an error starting the server: %v", err)
	}
}
