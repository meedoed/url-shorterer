package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/handlers"
	"github.com/meedoed/url-shorterer/pkg/logging"
)

const (
	getShortURL  = "/short"
	getSourceURL = "/source"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(app *fiber.App) {
	app.Post(getShortURL, h.GetShortURL)
	app.Get(getSourceURL, h.GetSourceURL)
}

func (h *handler) GetShortURL(c *fiber.Ctx) error {
	return c.SendString("http://shortURL")
}

func (h *handler) GetSourceURL(c *fiber.Ctx) error {
	return c.SendString("http://sourceURL")
}
