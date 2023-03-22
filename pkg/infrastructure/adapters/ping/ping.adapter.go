package adapters

import (
	"github.com/WildEgor/checker/pkg/config"
	domains "github.com/WildEgor/checker/pkg/infrastructure/domain"
)

type PingDto struct {
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

type IPingAdapter interface {
	GetApplicationStatus() *domains.Status
}

func NewPingAdapter(c *config.Config) *PingDto {
	return &PingDto{
		Version:     c.Version,
		Environment: c.GoEnv,
	}
}

func (service *PingDto) GetApplicationStatus() *domains.Status {
	return &domains.Status{
		Status:      "OK",
		Version:     service.Version,
		Environment: service.Environment,
	}
}
