package apigin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jeffotoni/gmelhorenvio/cmd/api.gin/handlers"
	"github.com/jeffotoni/gmelhorenvio/config"
	mw "github.com/jeffotoni/gmelhorenvio/internal/gin/middleware"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

func Run(dbLog *pg.DbConnection) {
	router := gin.Default()
	mw.Cors(router)
	mw.Logger(router)
	mw.Compress(router)

	srv := &http.Server{
		Addr:    config.SERVER_DOMAIN,
		Handler: router,
	}
	handlers.SetRoutes(router.Group("/v1"), dbLog)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
