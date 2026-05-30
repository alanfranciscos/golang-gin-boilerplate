package application

import (
	"testing"
)

func TestHealthCheckUseCase_GetHealth(t *testing.T) {
	version := "1.0.0"
	environment := "test"
	useCase := NewHealthCheckUseCase(version, environment)

	health := useCase.GetHealth()

	if health.Status != "UP" {
		t.Errorf("Expected status UP, got %s", health.Status)
	}
	if health.Version != version {
		t.Errorf("Expected version %s, got %s", version, health.Version)
	}
	if health.Environment != environment {
		t.Errorf("Expected environment %s, got %s", environment, health.Environment)
	}
	if health.Uptime == "" {
		t.Error("Expected uptime to be set")
	}
	if health.Timestamp.IsZero() {
		t.Error("Expected timestamp to be set")
	}
}
