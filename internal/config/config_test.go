package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set all required environment variables
	os.Setenv("APP_ENV", "production")
	os.Setenv("APP_PORT", "9090")
	os.Setenv("APP_VERSION", "2.0.0")
	os.Setenv("APP_NAME", "test-app")
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4318")
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", "grpc")

	cfg := LoadConfig()

	if cfg.AppEnv != "production" {
		t.Errorf("Expected APP_ENV production, got %s", cfg.AppEnv)
	}
	if cfg.AppPort != "9090" {
		t.Errorf("Expected APP_PORT 9090, got %s", cfg.AppPort)
	}
	if cfg.AppVersion != "2.0.0" {
		t.Errorf("Expected APP_VERSION 2.0.0, got %s", cfg.AppVersion)
	}
	if cfg.AppName != "test-app" {
		t.Errorf("Expected APP_NAME test-app, got %s", cfg.AppName)
	}
	if cfg.OTLPEndpoint != "localhost:4318" {
		t.Errorf("Expected OTLPEndpoint localhost:4318, got %s", cfg.OTLPEndpoint)
	}
	if cfg.OTLPProtocol != "grpc" {
		t.Errorf("Expected OTLPProtocol grpc, got %s", cfg.OTLPProtocol)
	}

	// Clean up
	os.Unsetenv("APP_ENV")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_VERSION")
	os.Unsetenv("APP_NAME")
	os.Unsetenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	os.Unsetenv("OTLPProtocol")
}
