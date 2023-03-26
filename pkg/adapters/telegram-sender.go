package adapters

import (
	"github.com/WildEgor/checker/pkg/config"
	domains "github.com/WildEgor/checker/pkg/domain"
)

type TelegramAdapter struct {
	tc *config.TelegramConfig
}

type ITelegramAdapter interface {
	GetApplicationStatus() *domains.StatusDomain
}

func NewTelegramAdapter(tc *config.TelegramConfig) *TelegramAdapter {
	return &TelegramAdapter{
		tc: tc,
	}
}

func (t *TelegramAdapter) SendAlert() {

}
