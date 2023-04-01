package config

import (
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type TelegramConfig struct {
	ChatId int64  `env:"TELEGRAM_CHAT_ID"`
	Token  string `env:"TELEGRAM_TOKEN"`
}

func NewTelegramConfig() *TelegramConfig {
	cfg := TelegramConfig{}

	if err := godotenv.Load(".env"); err == nil {
		if err := env.Parse(&cfg); err != nil {
			log.Printf("[TelegramConfig] %+v\n", err)
		}
	}

	return &cfg
}
