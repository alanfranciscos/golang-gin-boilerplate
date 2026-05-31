package web

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/alanf/go-boilerplate/internal/application/mocks"
	"github.com/alanf/go-boilerplate/internal/domain"
	"github.com/gin-gonic/gin"
)

func TestServer_HandleHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedHealth := domain.Health{
		Status:      "UP",
		Timestamp:   time.Now().UTC(),
		Version:     "1.0.0",
		Environment: "test",
		Uptime:      "1s",
	}

	mockHealthService := &mocks.HealthServiceMock{
		GetHealthFunc: func(ctx context.Context) domain.Health {
			return expectedHealth
		},
	}

	server := NewServer("8080", "test-app", mockHealthService)

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	server.router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	var response domain.Health
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if response.Status != expectedHealth.Status {
		t.Errorf("Expected status %s, got %s", expectedHealth.Status, response.Status)
	}
	if response.Version != expectedHealth.Version {
		t.Errorf("Expected version %s, got %s", expectedHealth.Version, response.Version)
	}
}

func TestServer_Run_Error(t *testing.T) {
	mockHealthService := &mocks.HealthServiceMock{}
	server := NewServer("invalid", "test-app", mockHealthService)

	err := server.Run()
	if err == nil {
		t.Error("Expected error when running server on invalid port, got nil")
	}
}
