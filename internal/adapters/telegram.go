package adapters

import (
	"strconv"

	"github.com/WildEgor/gChecker/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type ITelegramAdapter interface {
	Send(to string, text string) error
}

type TelegramAdapter struct {
	tc  *config.TelegramConfig
	bot *tgbotapi.BotAPI
}

func NewTelegramAdapter(tc *config.TelegramConfig) *TelegramAdapter {
	botTg, err := tgbotapi.NewBotAPI(tc.Token)
	if err != nil {
		log.Fatal("[TelegramAdapter] Init tgbotapi failed: ", err)
	}

	return &TelegramAdapter{
		tc:  tc,
		bot: botTg,
	}
}

func (t *TelegramAdapter) Send(to string, text string) error {
	log.Debug("[TelegramAdapter] Send message: ", to, text)
	var msg tgbotapi.MessageConfig

	if to == "" {
		msg = tgbotapi.NewMessage(t.tc.ChatId, text)
	} else {
		value, err := strconv.ParseInt(to, 64, 2)
		if err != nil {
			return err
		}
		msg = tgbotapi.NewMessage(value, text)
	}

	msg.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(msg)
	return err
}
