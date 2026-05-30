package telemetry

import (
	"context"
	"testing"

	"github.com/alanf/go-boilerplate/internal/config"
)

func TestInitOTel(t *testing.T) {
	cfg := &config.Config{
		AppName:      "test-app",
		OTLPProtocol: "http",
		OTLPEndpoint: "localhost:4318",
	}

	shutdown, err := InitOTel(context.Background(), cfg)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if shutdown == nil {
		t.Error("Expected shutdown function, got nil")
	}

	// Shutdown should work
	err = shutdown(context.Background())
	if err != nil {
		t.Errorf("Expected nil error on shutdown, got %v", err)
	}
}
