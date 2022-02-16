package config

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	App      *App
	Telegram *Telegram
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.Wrap(err, "error loading .env file")
	}

	app, err := LoadAppConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tg, err := LoadTelegramConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Config{
		App:      app,
		Telegram: tg,
	}, nil
}
