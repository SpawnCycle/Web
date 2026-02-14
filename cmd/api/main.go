package main

import (
	"context"
	"log"
	"smaash-web/internal/controllers"
	"smaash-web/internal/repository"
	"smaash-web/internal/server"
	"smaash-web/internal/services"
)

func main() {
	appContext := context.Background()
	userRepo := repository.NewGormUserRepo()
	playerProfileRepo := repository.NewGormPlayerProfileRepo()
	userStatsService := services.NewUserStatsService(userRepo)
	playerProfileService := services.NewPlayerProfileService(playerProfileRepo)
	authnService := services.NewAuthenticationService(userRepo)

	levelRepo := repository.NewGormLevelRepo()

	srv := server.NewServer(
		controllers.NewUserController(userStatsService),
		controllers.NewAuthnController(authnService),
		controllers.NewGameAuthController(authnService, playerProfileService, userRepo),
		controllers.NewLevelsController(levelRepo),
	).MountRoutes()

	if err := srv.Run(appContext); err != nil {
		log.Fatalf("There was an error starting the server: %v", err)
	}
}
