package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/config"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func UseCache(c *gin.Context) {
	if c.GetHeader("Cache-Control") == "no-cache" {
		c.Next()
		return
	}

	rawBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.NewHttpErr(err))
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	hash := cache.GetHash(rawBody)
	body, found := config.Cache.Get(hash)
	if !found {
		c.Next()
		return
	}

	c.Data(http.StatusOK, "application/json", body.([]byte))
	c.Abort()
}
