package config

import (
	"errors"
	"os"
)

type Telegram struct {
	Token string
}

func LoadTelegramConfig() (*Telegram, error) {
	token, _ := os.LookupEnv("TELEGRAM_TOKEN")
	if token == "" {
		return nil, errors.New("TELEGRAM_TOKEN is unset")
	}

	return &Telegram{
		Token: token,
	}, nil
}
