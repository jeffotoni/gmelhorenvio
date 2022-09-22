package middleware

import (
	"os"

	glogger "github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/config"
)

func Logger(r *gin.Engine) {
	if os.Getenv(config.ENV_AMBI) != "prod" {
		r.Use(glogger.SetLogger())
	}
	return
}
