package pkg

import (
	"github.com/WildEgor/checker/pkg/config"
	"github.com/WildEgor/checker/pkg/handlers"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
)

var AppSet = wire.NewSet(NewApp, handlers.Set, config.Set)

func NewApp(status *handlers.HealthCheckHandler) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
	}))

	v1 := app.Group("/api/v1")

	// Server endpoint - sanity check that the server is running
	statusGroup := v1.Group("/health")
	statusGroup.Get("/check", status.HealthCheckHandle)

	log.Info("Application is running on port...")
	return app
}
