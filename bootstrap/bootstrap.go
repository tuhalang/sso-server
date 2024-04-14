package bootstrap

import (
	"context"
	"fmt"
	"github.com/tuhalang/authen/api/grpc_server"
	"github.com/tuhalang/authen/api/rest/controller"
	"github.com/tuhalang/authen/api/rest/middleware"
	"github.com/tuhalang/authen/api/rest/route"
	"github.com/tuhalang/authen/config"
	"github.com/tuhalang/authen/internal/logger"
	"github.com/tuhalang/authen/repository"
	"github.com/tuhalang/authen/store"
	"github.com/tuhalang/authen/usecase"
)

type Application struct {
	appConfig *config.Config
	restRoute route.RestRoute
}

func Run(configFile string) {
	log := logger.Get()
	ctx := context.Background()
	appConfig := config.LoadConfig(configFile)

	appStores := make(map[string]store.Store)
	keyStores := make(map[string]config.KeyStore)

	for _, tenantCfg := range appConfig.Tenants {
		dbStore, err := store.NewStore(ctx, tenantCfg)
		if err != nil {
			log.Panic().Err(err).Msgf("Cannot init store %s", tenantCfg.TenantName)
		} else {
			appStores[tenantCfg.TenantName] = dbStore
		}
		defer dbStore.Close()

		keyStores[tenantCfg.TenantName] = config.KeyStore{
			JwtKey: tenantCfg.JwtKey,
			JwtExp: tenantCfg.JwtExpiredTime,
		}
	}

	userRepository := repository.NewUserRepository(appStores)
	sessionRepository := repository.NewSessionRepository(appStores)

	loginUseCase := usecase.NewLoginUseCase(userRepository, sessionRepository, keyStores)
	validationUseCase := usecase.NewValidationUseCase(keyStores)

	loginController := controller.NewLoginController(loginUseCase)
	validateController := controller.NewValidateController(validationUseCase)

	loggingMiddleware := middleware.LoggingMiddleware()

	restRoute := route.NewRestRoute(loginController, validateController, loggingMiddleware)

	serverAddress := fmt.Sprintf("%s:%d", appConfig.RestServer.Host, appConfig.RestServer.Port)
	go func() {
		restRoute.Run(serverAddress)
	}()

	grpcServer := grpc_server.NewGrpcServer(appConfig.GrpcServer, loginUseCase, validationUseCase)
	grpcServer.Start()
}
