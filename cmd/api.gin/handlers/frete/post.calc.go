package frete

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/internal/credentials/auth"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
	"github.com/jeffotoni/gmelhorenvio/repo/frete"
)

func PostCalc(dbLog *pg.DbConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.NewHttpErr(err))
			return
		}

		respBody, code, err := frete.Calc(body)
		if err != nil {
			c.JSON(code, models.NewHttpErr(err))
			return
		}

		if code == http.StatusUnauthorized {
			_, err := auth.RefreshToken()
			if err != nil {
				c.JSON(http.StatusInternalServerError, models.NewHttpErr(err))
				return
			}

			respBody, code, err = frete.Calc(body)
			if err != nil {
				c.JSON(code, models.NewHttpErr(err))
				return
			}
		}

		config.Cache.Set(cache.GetHash(body), respBody, 0)

		c.Status(code)
		c.Header("Content-Type", "application/json")
		c.Writer.Write(respBody)
	}
}
