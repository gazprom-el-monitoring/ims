package http

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Ping(c *fiber.Ctx) error {
	k := map[string]interface{}{
		"ok": true,
	}
	return c.JSON(k)
}
