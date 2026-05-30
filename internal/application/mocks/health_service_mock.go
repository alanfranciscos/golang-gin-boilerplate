package mocks

import (
	"github.com/alanf/go-boilerplate/internal/domain"
)

type HealthServiceMock struct {
	GetHealthFunc func() domain.Health
}

func (m *HealthServiceMock) GetHealth() domain.Health {
	return m.GetHealthFunc()
}
