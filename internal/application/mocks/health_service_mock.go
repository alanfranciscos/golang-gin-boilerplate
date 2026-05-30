package mocks

import (
	"context"
	"github.com/alanf/go-boilerplate/internal/domain"
)

type HealthServiceMock struct {
	GetHealthFunc func(ctx context.Context) domain.Health
}

func (m *HealthServiceMock) GetHealth(ctx context.Context) domain.Health {
	return m.GetHealthFunc(ctx)
}
