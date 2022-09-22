package cache

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/models"
	"github.com/jeffotoni/gmelhorenvio/repo/cache"
)

func DeleteKey(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, models.NewHttpErrStr("url param 'key' is required"))
		return
	}

	err := cache.DeleteKey(key)
	if err != nil {
		if err == cache.ErrCacheNotFound {
			c.JSON(http.StatusNotFound, models.NewHttpErr(err))
			return
		}
		c.JSON(http.StatusBadRequest, models.NewHttpErr(err))
		return
	}

	c.Status(http.StatusOK)
}
