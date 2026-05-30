package application

import (
	"context"
	"time"

	"github.com/alanf/go-boilerplate/internal/domain"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type healthCheckUseCase struct {
	version     string
	environment string
}

func NewHealthCheckUseCase(version, environment string) HealthService {
	return &healthCheckUseCase{
		version:     version,
		environment: environment,
	}
}

func (u *healthCheckUseCase) GetHealth(ctx context.Context) domain.Health {
	return domain.Health{
		Status:      "UP",
		Timestamp:   time.Now().UTC(),
		Version:     u.version,
		Environment: u.environment,
		Uptime:      time.Since(startTime).Truncate(time.Second).String(),
	}
}
