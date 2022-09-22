package headers

import "github.com/gofiber/fiber/v2"

func IP(c *fiber.Ctx) string {
	ipReal := string(c.Request().Header.Peek("X-Real-IP"))
	if len(ipReal) <= 0 {
		return "127.0.0.1"
	}
	return ipReal
}
