package adapters

import (
	"github.com/WildEgor/checker/pkg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
)

type TelegramAdapter struct {
	tc  *config.TelegramConfig
	bot *tgbotapi.BotAPI
}

type ITelegramAdapter interface {
	SendAlert(resource string, status string) error
}

func NewTelegramAdapter(tc *config.TelegramConfig) *TelegramAdapter {
	botTg, err := tgbotapi.NewBotAPI(tc.Token)
	if err != nil {
		log.Fatal("Token not provided", err)
	}

	return &TelegramAdapter{
		tc:  tc,
		bot: botTg,
	}
}

func (t *TelegramAdapter) SendAlert(r string, s string) error {
	log.Info("Send alert!")
	msg := tgbotapi.NewMessage(t.tc.ChatId, "Service <code>"+r+"</code> is down\n"+"Status: <b>"+s+"</b>")
	msg.ParseMode = tgbotapi.ModeHTML
	_, err := t.bot.Send(msg)
	return err
}
