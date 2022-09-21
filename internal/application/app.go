package application

import (
	"context"
	"tag-value-finder/internal/adapters/config"
	"tag-value-finder/internal/adapters/http2"
	"tag-value-finder/internal/adapters/rmq"
	"tag-value-finder/internal/ports"

	"github.com/rs/zerolog/log"
)

var yawmRMQ ports.MQ //nolint:gochecknoglobals

func Start(ctx context.Context) {
	log.Debug().Msg("Read config")
	cfg, _ := config.NewConfig()
	log.Debug().Msg("Starting application")
	_ = http2.HealthCheckHandler(ctx, cfg.HealthCheckAddr)
	yawmRMQ, _ = rmq.NewYawm(ctx, cfg.MQConnURI, cfg.InQueryName, cfg.OutQueryName)
	_ = yawmRMQ.LaunchConsumer()
}

func Stop(ctx context.Context) {
	_ = yawmRMQ.Disconnect()
	log.Debug().Msg("Application stopped")
}
