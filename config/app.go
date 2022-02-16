package config

import (
	"errors"
	"os"
	"strings"
	"time"
)

type App struct {
	// Account name
	Account string

	// Password
	Password string

	// How long before to inform user there is prompt items
	PriorTime time.Duration

	// How often to fetch latest data
	UpdateInterval time.Duration

	Debug bool
}

func LoadAppConfig() (*App, error) {
	var (
		account   string
		password  string
		priorTime = 5 * time.Minute
		interval  = 10 * time.Minute
		debug     bool
	)

	account, _ = os.LookupEnv("ACCOUNT")
	if account == "" {
		return nil, errors.New("ACCOUNT is unset")
	}

	password, _ = os.LookupEnv("PASSWORD")
	if password == "" {
		return nil, errors.New("PASSWORD is unset")
	}

	debugVal, _ := os.LookupEnv("APP_DEBUG")
	debugVal = strings.TrimSpace(debugVal)
	if debugVal != "" && debugVal != "false" && debugVal != "0" {
		debug = true
	} else {
		debug = false
	}

	return &App{
		Account:        account,
		Password:       password,
		PriorTime:      priorTime,
		UpdateInterval: interval,
		Debug:          debug,
	}, nil
}
