package common

import "github.com/gofiber/fiber/v2"

var (
	header = []byte("Strict-Transport-Security")
	value  = []byte("max-age=63072000; includeSubDomains; preload")
)

// HTSTMiddleware adds the HTST Header
// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
func HTSTMiddleware(c *fiber.Ctx) error {
	c.Response().Header.SetCanonical(header, value)
	return nil
}
