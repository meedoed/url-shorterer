package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/user"
	"github.com/meedoed/url-shorterer/pkg/logging"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("start app")
	app := fiber.New()

	handler := user.NewHandler(logger)
	handler.Register(app)

	start(app)
}

func start(app *fiber.App) {
	logger := logging.GetLogger()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!")
	})

	logger.Info("Server is listening...")
	logger.Fatal(app.Listen(":8000"))
}
