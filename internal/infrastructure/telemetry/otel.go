package telemetry

import (
	"context"
	"os"

	"github.com/alanf/go-boilerplate/internal/config"
	otelcol "github.com/alanfranciscos/otel-collector/pkg/telemetry"
)

func InitOTel(ctx context.Context, cfg *config.Config) (func(context.Context) error, error) {
	os.Setenv("OTEL_SERVICE_NAME", cfg.AppName)
	os.Setenv("OTEL_EXPORTER_OTLP_PROTOCOL", cfg.OTLPProtocol)
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", cfg.OTLPEndpoint)

	shutdown, err := otelcol.NewTelemetry(&cfg.AppName).Initialize(ctx)
	if err != nil {
		return nil, err
	}

	return shutdown, nil
}
