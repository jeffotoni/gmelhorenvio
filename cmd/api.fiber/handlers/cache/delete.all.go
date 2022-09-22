package cache

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func DeleteAll(c *fiber.Ctx) error {
	cache.DeleteAll()
	return c.SendStatus(http.StatusOK)
}
