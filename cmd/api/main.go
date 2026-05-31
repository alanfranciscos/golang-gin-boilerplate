package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alanf/go-boilerplate/internal/application"
	"github.com/alanf/go-boilerplate/internal/config"
	"github.com/alanf/go-boilerplate/internal/infrastructure/telemetry"
	"github.com/alanf/go-boilerplate/internal/infrastructure/web"
)

var logFatal = log.Fatalf

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := Run(ctx); err != nil {
		logFatal("Failed to start server: %v", err)
	}
}

func Run(ctx context.Context) error {
	cfg := config.LoadConfig()

	// Initialize OTel
	shutdown, err := telemetry.InitOTel(ctx, cfg)
	if err != nil {
		return err
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := shutdown(shutdownCtx); err != nil {
			log.Printf("Failed to shutdown telemetry: %v", err)
		}
	}()

	healthService := application.NewHealthCheckUseCase(cfg.AppVersion, cfg.AppEnv)
	server := web.NewServer(cfg.AppPort, cfg.AppName, healthService)

	log.Printf("Starting server on port %s", cfg.AppPort)
	
	// Start server in a goroutine
	errChan := make(chan error, 1)
	go func() {
		if err := server.Run(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		log.Println("Shutting down server...")
	}

	return nil
}
