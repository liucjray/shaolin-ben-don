package app

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/wolftotem4/shaolin-ben-don/config"
	"github.com/wolftotem4/shaolin-ben-don/internal/client"
	_ "modernc.org/sqlite"
)

type App struct {
	Config     *config.Config
	DB         *sqlx.DB
	RealClient *client.HttpClient
	Client     client.Client
	Bot        *tgbotapi.BotAPI
}

func Register() (*App, error) {
	var (
		cfig    *config.Config
		db      *sqlx.DB
		real    *client.HttpClient
		client_ client.Client
		err     error
	)

	cfig, err = config.LoadConfig()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	db, err = sqlx.Open("sqlite", "db.sqlite")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	real, err = client.NewClient(client.NewDatabaseStore(db))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if cfig.App.Debug {
		client_ = client.WrapLogger(real, "logs")
	} else {
		client_ = real
	}

	bot, err := createTelegramBotClient(cfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &App{
		Config:     cfig,
		DB:         db,
		RealClient: real,
		Client:     client_,
		Bot:        bot,
	}, nil
}

func createTelegramBotClient(config *config.Config) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.Telegram.Token)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bot.Debug = config.App.Debug

	return bot, nil
}
