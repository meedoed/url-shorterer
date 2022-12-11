package urlcollection

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/meedoed/url-shorterer/internal/handlers"
	"github.com/meedoed/url-shorterer/pkg/logging"
)

const (
	getShortURL  = "/short"
	getSourceURL = "/source"
)

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(app *fiber.App) {
	app.Post(getShortURL, h.GetShortURL)
	app.Get(getSourceURL, h.GetSourceURL)
}

func (h *handler) GetShortURL(c *fiber.Ctx) error {
	url := URL{}
	str := c.Body()
	if err := json.Unmarshal(str, &url); err != nil {
		return err
	}

	shortURL, err := h.repository.Create(c.Context(), url)
	if err != nil {
		return err
	}
	c.SendString(shortURL)
	return nil
}

func (h *handler) GetSourceURL(c *fiber.Ctx) error {
	url := c.Query("short")
	sourceURL, err := h.repository.Find(c.Context(), url)
	if err != nil {
		return err
	}
	c.SendString(sourceURL)
	return nil
}
