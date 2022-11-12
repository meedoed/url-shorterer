package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/user"
	"log"
)

func main() {
	app := fiber.New()

	handler := user.NewHandler()
	handler.Register(app)

	start(app)
}

func start(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from fiber!")
	})

	log.Fatal(app.Listen(":8000"))
}
