package handlers

import (
	adapters "github.com/WildEgor/checker/pkg/adapters"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	PingAdapter adapters.IPingAdapter
}

func NewHealthCheckHandler(pingAdapter adapters.IPingAdapter) *HealthCheckHandler {
	return &HealthCheckHandler{
		PingAdapter: pingAdapter,
	}
}

func (s *HealthCheckHandler) Handle(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"isOk": true,
		"data": s.PingAdapter.GetApplicationStatus(),
	})
	return nil
}
