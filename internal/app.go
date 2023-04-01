package pkg

import (
	"context"
	"os"

	"github.com/WildEgor/gChecker/internal/adapters"
	"github.com/WildEgor/gChecker/internal/config"
	error_handler "github.com/WildEgor/gChecker/internal/errors"
	"github.com/WildEgor/gChecker/internal/router"
	"github.com/WildEgor/gChecker/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/google/wire"
	log "github.com/sirupsen/logrus"
)

var AppSet = wire.NewSet(
	NewApp,
	adapters.AdaptersSet,
	config.ConfigsSet,
	router.RouterSet,
	services.ServicesSet,
)

func NewApp(
	appConfig *config.AppConfig,
	router *router.Router,
	services *services.CheckerService,
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

	// Set logging settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	if !appConfig.IsProduction() {
		// HINT: some extra setting
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	router.Setup(app)

	// go services.Check()
	go services.ServicesCheck(context.Background())

	log.Info("Application is running on port...")
	return app
}
