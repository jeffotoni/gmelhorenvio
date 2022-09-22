package cache

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func DeleteKey(c *fiber.Ctx) error {
	key := c.Params("key")
	if key == "" {
		return c.Status(http.StatusBadRequest).JSON(models.NewHttpErrStr("url param 'key' is required"))
	}

	err := cache.DeleteKey(key)
	if err != nil {
		if err == cache.ErrCacheNotFound {
			return c.Status(http.StatusNotFound).JSON(models.NewHttpErr(err))
		}
		return c.Status(http.StatusBadRequest).JSON(models.NewHttpErr(err))
	}

	return c.SendStatus(http.StatusOK)
}
