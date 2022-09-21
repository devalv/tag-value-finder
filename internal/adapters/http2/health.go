package http2

import (
	"context"
	"net/http"
	"tag-value-finder/internal/domain/errors"
	"time"

	"github.com/rs/zerolog/log"
)

func HealthCheckHandler(ctx context.Context, addr string) error {
	log.Debug().Msgf("Trying to start http server on %s", addr)

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("ok"))
		})
		srv := &http.Server{
			ReadHeaderTimeout: 5 * time.Second,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
			Addr:              addr,
		}

		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal().Msgf(errors.HTTPHealthListenError, err)
		}
	}()

	log.Debug().Msgf("HealthCheckHandler started.")
	return nil
}
