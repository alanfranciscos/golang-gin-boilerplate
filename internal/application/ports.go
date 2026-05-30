package application

import (
	"context"
	"github.com/alanf/go-boilerplate/internal/domain"
)

type HealthService interface {
	GetHealth(ctx context.Context) domain.Health
}
