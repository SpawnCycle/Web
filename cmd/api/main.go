package main

import (
	"context"
	"log"
	"smaash-web/internal/repository"
	"smaash-web/internal/server"
)

func main() {
	appContext := context.Background()
	repoService := repository.NewRepositoryService(
		repository.NewGormUserRepo(),
	)

	if err := server.NewServer(&repoService).MountRoutes().Run(appContext); err != nil {
		log.Fatalf("There was an error starting the server: %v", err)
	}
}
