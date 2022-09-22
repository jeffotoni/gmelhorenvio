package healthz

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Healthz(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
