package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv            string
	AppPort           string
	AppVersion        string
	AppName           string
	OTLPEndpoint      string
	OTLPProtocol      string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	cfg := &Config{
		AppEnv:            requireEnv("APP_ENV"),
		AppPort:           requireEnv("APP_PORT"),
		AppVersion:        requireEnv("APP_VERSION"),
		AppName:           requireEnv("APP_NAME"),
		OTLPEndpoint:      requireEnv("OTEL_EXPORTER_OTLP_ENDPOINT"),
		OTLPProtocol:      requireEnv("OTEL_EXPORTER_OTLP_PROTOCOL"),
	}

	return cfg
}

func requireEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Required environment variable %s is not set", key)
	}
	return value
}
