package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/jeffotoni/gmelhorenvio/config"
)

func main() {
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(`{"ok": true}`)
	})

	app.Listen(":" + config.API_TEST_PORT)
}
