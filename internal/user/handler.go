package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/handlers"
)

const (
	getShortURL  = "/short"
	getSourceURL = "/source"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(app *fiber.App) {
	app.Post(getShortURL, GetShortURL)
	app.Get(getSourceURL, GetSourceURL)
}

func GetShortURL(c *fiber.Ctx) error {
	return c.SendString("http://shortURL")
}

func GetSourceURL(c *fiber.Ctx) error {
	return c.SendString("http://sourceURL")
}
