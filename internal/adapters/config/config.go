package config

import (
	"github.com/devalv/tag-value-finder/internal/domain/errors"
	"github.com/devalv/tag-value-finder/internal/domain/models"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog/log"
)

func NewConfig() (*models.Config, error) {
	var cfg models.Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatal().Msgf(errors.ConfigError, err)
	}

	return &cfg, nil
}
