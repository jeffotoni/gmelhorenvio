package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.gin/handlers/cache"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.gin/handlers/frete"
	"github.com/jeffotoni/gmelhorenvio/config"
	mw "github.com/jeffotoni/gmelhorenvio/internal/gin/middleware"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
	"github.com/ulule/limiter/v3"
	mwgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func SetRoutes(r *gin.RouterGroup, dbLog *pg.DbConnection) {
	lim := mwgin.NewMiddleware(limiter.New(
		memory.NewStore(),
		limiter.Rate{
			Period: config.LIMITER_EXPIRATION_SEC,
			Limit:  int64(config.LIMITER_MAX_REQUESTS),
		},
	))

	r.Use(lim, mw.ValidateStaticToken)

	cacheG := r.Group("/cache")
	cacheG.DELETE("/", cache.DeleteAll)
	cacheG.DELETE("/:key", cache.DeleteKey)

	freteG := r.Group("/frete", mw.UseCache)
	freteG.POST("/calc", frete.PostCalc(dbLog))
}
