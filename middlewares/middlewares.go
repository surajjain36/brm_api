package middlewares

import (
	"brm_api/utils/response"

	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 page
func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(response.HTTPResponse{
		Success: false,
		Message: "route not found",
		Data:    nil,
	})
}
