package frete

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
	"github.com/jeffotoni/gmelhorenvio/repo/frete"
)

func PostCalc(c *fiber.Ctx) error {
	respBody, code, err := frete.Calc(c.Body())
	if err != nil {
		return c.Status(code).JSON(models.NewHttpErr(err))
	}

	println("code:", code)
	if code == http.StatusUnauthorized {
		_, err := auth.RefreshToken()
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(models.NewHttpErr(err))
		}

		respBody, code, err = frete.Calc(c.Body())
		if err != nil {
			return c.Status(code).JSON(models.NewHttpErr(err))
		}
	}

	config.Cache.Set(cache.GetHash(c.Body()), respBody, 0)
	c.Response().Header.Add("Content-Type", "application/json")
	return c.Status(code).Send(respBody)
}
