package telemetry

import (
	"context"
	"os"

	"github.com/alanf/go-boilerplate/internal/config"
	"github.com/alanfranciscos/otel-lgtm-sdk-go/pkg/telemetry"
)

func InitOTel(ctx context.Context, cfg *config.Config) (func(context.Context) error, error) {
	os.Setenv("OTEL_SERVICE_NAME", cfg.AppName)
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", cfg.OTLPEndpoint)

	if err := telemetry.Init(ctx); err != nil {
		return nil, err
	}

	return func(ctx context.Context) error {
		return telemetry.Shutdown(ctx)
	}, nil
}
