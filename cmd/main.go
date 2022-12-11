package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/config"
	"github.com/meedoed/url-shorterer/internal/urlcollection"
	"github.com/meedoed/url-shorterer/internal/urlcollection/db"
	"github.com/meedoed/url-shorterer/pkg/client/postgresql"
	"github.com/meedoed/url-shorterer/pkg/logging"
	"log"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create app")
	app := fiber.New()

	cfg := config.GetConfig()
	postgres, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err != nil {
		log.Fatal(err)
	}
	repository := db.NewRepository(postgres, logger)
	handler := urlcollection.NewHandler(repository, logger)
	handler.Register(app)

	start(app, cfg)
}

func start(app *fiber.App, cfg *config.Config) {
	logger := logging.GetLogger()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!\n")

	})

	logger.Infof("Server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	logger.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port)))
}
