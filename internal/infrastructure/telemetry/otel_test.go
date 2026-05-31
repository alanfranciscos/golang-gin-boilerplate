package telemetry

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/alanf/go-boilerplate/internal/config"
)

func TestInitOTel(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("APP_VERSION", "1.0.0")
	os.Setenv("APP_NAME", "test-app")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4317")

	cfg := &config.Config{
		AppName:      "test-app",
		OTLPEndpoint: "localhost:4317",
	}

	shutdown, err := InitOTel(context.Background(), cfg)
	if err != nil {
		t.Fatalf("Expected nil error on InitOTel, got %v", err)
	}

	if shutdown == nil {
		t.Fatal("Expected shutdown function, got nil")
	}

	// In test environments without a collector, Shutdown might return an error
	// due to failure to flush metrics/traces. We call it to ensure it doesn't panic.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	_ = shutdown(ctx)
}
