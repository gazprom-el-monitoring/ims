package http

import (
	"github.com/gazprom-el-monitoring/ims/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) Init(app *fiber.App) {
	app.Get("/ping", h.Ping)
}
