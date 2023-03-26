package config

import (
	"github.com/spf13/viper"
)

// HINT: holds on telegram settings
type TelegramConfig struct {
	ChatId string `dotenv:"TELEGRAM_CHAT_ID"`
	Token  string `dotenv:"TELEGRAM_TOKEN"`
}

func NewTelegramConfig() *TelegramConfig {
	var config TelegramConfig

	viper.Unmarshal(&config)

	return &config
}
