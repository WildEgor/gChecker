package handlers

import (
	adapters "github.com/WildEgor/checker/pkg/infrastructure/adapters/ping"
	"github.com/gofiber/fiber"
)

type HealthCheckHandler struct {
	PingAdapter adapters.IPingAdapter
}

func NewHealthCheckHandler(pingAdapter adapters.IPingAdapter) *HealthCheckHandler {
	return &HealthCheckHandler{
		PingAdapter: pingAdapter,
	}
}

func (s *HealthCheckHandler) HealthCheckHandle(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"isOk": true,
		"data": s.PingAdapter.GetApplicationStatus(),
	})
	return
}
