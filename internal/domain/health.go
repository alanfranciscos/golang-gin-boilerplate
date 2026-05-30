package domain

import "time"

type Health struct {
	Status      string    `json:"status"`
	Timestamp   time.Time `json:"timestamp"`
	Version     string    `json:"version"`
	Environment string    `json:"environment"`
	Uptime      string    `json:"uptime"`
}
