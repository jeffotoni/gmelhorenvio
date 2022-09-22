package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func UseCache(c *fiber.Ctx) error {
	if c.Get("Cache-Control") == "no-cache" {
		return c.Next()
	}

	hash := cache.GetHash(c.Body())
	// log.Println(hash)
	body, found := config.Cache.Get(hash)
	if !found {
		return c.Next()
	}

	c.Response().Header.Add("Content-Type", "application/json")
	return c.Status(http.StatusOK).Send(body.([]byte))
}
