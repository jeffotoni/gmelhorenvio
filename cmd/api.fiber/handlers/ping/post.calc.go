package ping

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetPing(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
