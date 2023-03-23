package adapters

import (
	"github.com/WildEgor/checker/pkg/config"
	domains "github.com/WildEgor/checker/pkg/infrastructure/domain"
)

type PingAdapter struct {
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

type IPingAdapter interface {
	GetApplicationStatus() *domains.Status
}

func NewPingAdapter(c *config.Config) *PingAdapter {
	return &PingAdapter{
		Version:     c.Version,
		Environment: c.GoEnv,
	}
}

func (service *PingAdapter) GetApplicationStatus() *domains.Status {
	return &domains.Status{
		Status:      "ok",
		Version:     service.Version,
		Environment: service.Environment,
	}
}
