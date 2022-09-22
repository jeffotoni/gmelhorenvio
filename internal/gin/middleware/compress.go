package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Compress(r *gin.Engine) {
	r.Use(gzip.Gzip(gzip.DefaultCompression))
}
