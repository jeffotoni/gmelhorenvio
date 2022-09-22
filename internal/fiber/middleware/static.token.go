package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/config"
)

func ValidateStaticToken(c *fiber.Ctx) error {
	if c.Get("Authorization") != config.API_STATIC_TOKEN {
		return c.SendStatus(http.StatusUnauthorized)
	}
	return c.Next()
}
