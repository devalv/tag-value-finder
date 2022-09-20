package application

import (
	"context"
	"tag-value-finder/internal/adapters/http2"
	"tag-value-finder/internal/adapters/rmq"
	"tag-value-finder/internal/ports"

	"github.com/rs/zerolog/log"
)

var yawmRMQ ports.MQ //nolint:gochecknoglobals

func Start(ctx context.Context) {
	log.Debug().Msg("Starting application")
	_ = http2.HealthCheckHandler(ctx, ":3333")
	yawmRMQ, _ = rmq.NewYawm(ctx, "amqp://guest:guest@localhost:5672/", "biba", "boba")
	_ = yawmRMQ.LaunchConsumer()
}

func Stop(ctx context.Context) {
	_ = yawmRMQ.Disconnect()
	log.Debug().Msg("Application stopped")
}
