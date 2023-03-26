package pkg

import (
	"github.com/WildEgor/checker/pkg/config"
	error_handler "github.com/WildEgor/checker/pkg/errors"
	"github.com/WildEgor/checker/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
)

var AppSet = wire.NewSet(NewApp, config.ConfigsSet, router.RouterSet)

func NewApp(
	appConfig *config.AppConfig,
	router *router.Router,
) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: error_handler.ErrorHandler,
	})

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin, Content-Type, Accept, Content-Length, Accept-Language, Accept-Encoding, Connection, Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	app.Use(recover.New())

	if !appConfig.IsProduction() {
		// TODO: add swagger here
	}

	router.Setup(app)

	log.Info("Application is running on port...")
	return app
}
