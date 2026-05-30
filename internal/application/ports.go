package application

import "github.com/alanf/go-boilerplate/internal/domain"

type HealthService interface {
	GetHealth() domain.Health
}
