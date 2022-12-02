package frete

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
	"github.com/jeffotoni/gmelhorenvio/internal/log"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
	"github.com/jeffotoni/gmelhorenvio/repo/frete"
)

func PostCalc(dbLog *pg.DbConnection) fiber.Handler {
	return func(c *fiber.Ctx) error {
		service := c.Get("X-Ecommerce")
		if service == "" {
			return c.Status(http.StatusBadRequest).JSON(models.NewHttpErrStr("must provide 'Ecommerce' header"))
		}

		respBody, code, err := frete.Calc(c.Body())
		if err != nil {
			log.InsertLogMelhorEnvio(dbLog, log.ErrStatus, service, "error frete.Calc", string(c.Body()))
			return c.Status(code).JSON(models.NewHttpErr(err))
		}

		// println("code:", code)
		if code == http.StatusUnauthorized {
			_, err := auth.RefreshToken()
			if err != nil {
				log.InsertLogMelhorEnvio(dbLog, log.ErrStatus, service, "error auth.RefreshToken", string(c.Body()))
				return c.Status(http.StatusInternalServerError).JSON(models.NewHttpErr(err))
			}

			respBody, code, err = frete.Calc(c.Body())
			if err != nil {
				log.InsertLogMelhorEnvio(dbLog, log.ErrStatus, service, "error frete.Calc", string(c.Body()))
				return c.Status(code).JSON(models.NewHttpErr(err))
			}
		}

		config.Cache.Set(cache.GetHash(c.Body()), respBody, 0)
		c.Response().Header.Add("Content-Type", "application/json")
		return c.Status(code).Send(respBody)
	}
}
