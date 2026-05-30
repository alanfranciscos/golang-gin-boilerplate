package telemetry

import (
	"context"
	"os"
	"testing"

	"github.com/alanf/go-boilerplate/internal/config"
)

func TestInitOTel(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("APP_NAME", "test-app")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4318")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "http")

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
