package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.fiber/handlers/cache"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.fiber/handlers/frete"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.fiber/handlers/healthz"
	"github.com/jeffotoni/gmelhorenvio/cmd/api.fiber/handlers/ping"
	"github.com/jeffotoni/gmelhorenvio/config"
	hd "github.com/jeffotoni/gmelhorenvio/internal/fiber/headers"
	mw "github.com/jeffotoni/gmelhorenvio/internal/fiber/middleware"
	pg "github.com/jeffotoni/gmelhorenvio/internal/psql"
)

func SetRoutes(r fiber.Router, dbLog *pg.DbConnection) {
	lim := limiter.New(limiter.Config{
		Next:       nil,
		Max:        config.LIMITER_MAX_REQUESTS,
		Expiration: config.LIMITER_EXPIRATION_SEC,
		KeyGenerator: func(c *fiber.Ctx) string {
			return hd.IP(c)
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString(`{"msg":"Much Request #bloqued"}`)
		},
	})

	r.Use(lim, mw.ValidateStaticToken)
	r.Get("/healthz", healthz.Healthz)
	r.Get("/ping", ping.GetPing)

	cacheG := r.Group("/cache")
	cacheG.Delete("/", cache.DeleteAll)
	cacheG.Delete("/:key", cache.DeleteKey)

	// go-cache need to implement in config cache
	freteG := r.Group("/frete", mw.UseCache)
	freteG.Post("/calc", frete.PostCalc(dbLog))
}
