package http

import (
	"github.com/gazprom-el-monitoring/ims/internal/models"
	"github.com/gazprom-el-monitoring/ims/pkg/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	claims, err := auth.ParseToken(token)

	if err != nil || claims == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	role, err := models.TryParseRole((*claims)["role"].(string))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	c.Locals("role", role)
	return c.Next()
}
