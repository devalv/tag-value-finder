package application

import (
	"context"
	"tag-value-finder/internal/domain/crawler"

	"github.com/rs/zerolog/log"
)

func Start(ctx context.Context) {
	log.Debug().Msg("Starting application")
	log.Debug().Msg("Application started")

	q := crawler.GetH1("https://www.labirint.ru/books/622004/")
	log.Debug().Msg(q)
}

func Stop(ctx context.Context) {
	log.Debug().Msg("Stopping application")
	log.Debug().Msg("Application stopped")
}
