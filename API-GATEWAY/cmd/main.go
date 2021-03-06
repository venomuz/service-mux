package main

import (
	"github.com/venomuz/project5/API-GATEWAY/api"
	"github.com/venomuz/project5/API-GATEWAY/config"
	"github.com/venomuz/project5/API-GATEWAY/pkg/logger"
	"github.com/venomuz/project5/API-GATEWAY/services"
	"net/http"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := http.ListenAndServe(cfg.HTTPPort, server); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
