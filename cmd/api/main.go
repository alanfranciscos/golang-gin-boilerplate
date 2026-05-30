package main

import (
	"log"

	"github.com/alanf/go-boilerplate/internal/application"
	"github.com/alanf/go-boilerplate/internal/config"
	"github.com/alanf/go-boilerplate/internal/infrastructure/web"
)

var logFatal = log.Fatalf

func main() {
	if err := Run(); err != nil {
		logFatal("Failed to start server: %v", err)
	}
}

func Run() error {
	cfg := config.LoadConfig()

	healthService := application.NewHealthCheckUseCase(cfg.AppVersion, cfg.AppEnv)
	server := web.NewServer(cfg.AppPort, healthService)

	log.Printf("Starting server on port %s", cfg.AppPort)
	return server.Run()
}
