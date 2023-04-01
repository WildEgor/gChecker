package handlers

import (
	adapters "github.com/WildEgor/gChecker/internal/adapters"
	"github.com/gofiber/fiber/v2"
)

type HealthCheckHandler struct {
	pa adapters.IPingAdapter
}

func NewHealthCheckHandler(pa adapters.IPingAdapter) *HealthCheckHandler {
	return &HealthCheckHandler{
		pa: pa,
	}
}

func (s *HealthCheckHandler) Handle(c *fiber.Ctx) error {
	c.JSON(fiber.Map{
		"isOk": true,
		"data": s.pa.GetApplicationStatus(),
	})
	return nil
}
