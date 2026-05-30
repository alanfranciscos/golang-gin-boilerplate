package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test with environment variables
	os.Setenv("APP_ENV", "production")
	os.Setenv("APP_PORT", "9090")
	os.Setenv("APP_VERSION", "2.0.0")

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

	// Test default values
	os.Unsetenv("APP_ENV")
	os.Unsetenv("APP_PORT")
	os.Unsetenv("APP_VERSION")

	cfg = LoadConfig()

	if cfg.AppEnv != "development" {
		t.Errorf("Expected APP_ENV development, got %s", cfg.AppEnv)
	}
	if cfg.AppPort != "8080" {
		t.Errorf("Expected APP_PORT 8080, got %s", cfg.AppPort)
	}
	if cfg.AppVersion != "1.0.0" {
		t.Errorf("Expected APP_VERSION 1.0.0, got %s", cfg.AppVersion)
	}
}
