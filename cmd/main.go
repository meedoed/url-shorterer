package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/config"
	"github.com/meedoed/url-shorterer/internal/user"
	"github.com/meedoed/url-shorterer/pkg/logging"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create app")
	app := fiber.New()

	cfg := config.GetConfig()

	handler := user.NewHandler(logger)
	handler.Register(app)

	start(app, cfg)
}

func start(app *fiber.App, cfg *config.Config) {
	logger := logging.GetLogger()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!")
	})

	logger.Infof("Server is listening port %s:%s", cfg.Listen.BindIp, cfg.Listen.Port)
	logger.Fatal(app.Listen(fmt.Sprintf("%s:%s", cfg.Listen.BindIp, cfg.Listen.Port)))
}
