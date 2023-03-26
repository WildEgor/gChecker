package adapters

import (
	"github.com/WildEgor/checker/pkg/config"
	domains "github.com/WildEgor/checker/pkg/domain"
)

type PingAdapter struct {
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

type IPingAdapter interface {
	GetApplicationStatus() *domains.StatusDomain
}

func NewPingAdapter(c *config.AppConfig) *PingAdapter {
	return &PingAdapter{
		Version:     c.Version,
		Environment: c.GoEnv,
	}
}

func (service *PingAdapter) GetApplicationStatus() *domains.StatusDomain {

	// TODO: check system health here

	return &domains.StatusDomain{
		Status:      "ok",
		Version:     service.Version,
		Environment: service.Environment,
	}
}
