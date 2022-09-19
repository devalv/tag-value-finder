package application

import (
	"context"
	"tag-value-finder/internal/adapters/rmq"

	"github.com/rs/zerolog/log"
)

var yawmRMQ *rmq.YawmRmq //nolint:gochecknoglobals

func Start(ctx context.Context) {
	log.Debug().Msg("Starting application")
	yawmRMQ, _ = rmq.NewYawm(ctx, "amqp://guest:guest@localhost:5672/", "biba", "boba")
	_ = yawmRMQ.LaunchConsumer()
	log.Debug().Msg("Application started")
}

func Stop(ctx context.Context) {
	log.Debug().Msg("Stopping application")
	_ = yawmRMQ.Disconnect()
	log.Debug().Msg("Application stopped")
}
