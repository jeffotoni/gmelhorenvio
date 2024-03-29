package apifiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.fiber/handlers"
	"github.com/jeffotoni/gmelhorenvio/config"
	mw "github.com/jeffotoni/gmelhorenvio/internal/fiber/middleware"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

func Run(dbLog *pg.DbConnection) {
	app := fiber.New(
		fiber.Config{
			BodyLimit: 1024 * 1024,
			Prefork:   false,
		},
	)
	mw.Cors(app)
	mw.Logger(app)
	mw.Compress(app)
	handlers.SetRoutes(app.Group("/v1"), dbLog)
	app.Listen(config.SERVER_DOMAIN)
}
