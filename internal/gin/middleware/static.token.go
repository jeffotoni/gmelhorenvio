package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/config"
)

func ValidateStaticToken(c *gin.Context) {
	if c.GetHeader("Authorization") != config.API_STATIC_TOKEN {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Next()
}
