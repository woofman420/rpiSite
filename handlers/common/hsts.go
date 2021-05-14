package common

import "github.com/gofiber/fiber/v2"

var (
	header = []byte("Strict-Transport-Security")
	value  = []byte("max-age=63072000; includeSubDomains; preload")
)

// HSTSMiddleware adds the HSTS Header
// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
func HSTSMiddleware(c *fiber.Ctx) error {
	c.Response().Header.SetCanonical(header, value)
	return c.Next()
}
