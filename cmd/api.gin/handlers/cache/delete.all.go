package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func DeleteAll(c *gin.Context) {
	cache.DeleteAll()
	c.Status(http.StatusOK)
}
