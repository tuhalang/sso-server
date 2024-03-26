package bootstrap

import (
	"context"
	"github.com/tuhalang/authen/config"
	"github.com/tuhalang/authen/internal/logger"
	store "github.com/tuhalang/authen/store"
)

type Application struct{}

func App(configFile string) *Application {
	log := logger.Get()
	ctx := context.Background()
	appConfig := config.LoadConfig(configFile)

	appStores := make(map[string]store.Store)
	for _, dbConfig := range appConfig.Databases {
		dbStore, err := store.NewStore(ctx, dbConfig)
		if err != nil {
			log.Fatal().Err(err).Msgf("Cannot init store %s", dbConfig.SpaceName)
		} else {
			appStores[dbConfig.SpaceName] = dbStore
		}
	}

	return &Application{}
}
