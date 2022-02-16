package config

type Telegram struct {
	Token string
}

func LoadTelegramConfig() *Telegram {
	return &Telegram{
		Token: "5274056807:AAFCd9qf2eWtnQGTph1x7IprMUTDhb9IhNk",
	}
}
