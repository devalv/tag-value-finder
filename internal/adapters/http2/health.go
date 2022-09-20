package http2

import (
	"context"
	"net/http"
	"tag-value-finder/internal/domain/errors"

	"github.com/rs/zerolog/log"
)

func HealthCheckHandler(ctx context.Context, addr string) error {
	log.Debug().Msgf("Trying to start http server on %s", addr)

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal().Msgf(errors.HTTPHealthListenError, err)
		}
	}()

	log.Debug().Msgf("HealthCheckHandler started.")
	return nil
}
