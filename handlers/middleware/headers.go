package middleware

import "github.com/gofiber/fiber/v2"

var (
	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Alt-Svc
	altSVCHeader = []byte("Alt-Svc")
	altSVCValue  = []byte("h2=\":443\"")

	// See: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security
	hstsHeader = []byte("Strict-Transport-Security")
	hstsValue  = []byte("max-age=63072000; includeSubDomains; preload")
)

// AddHeaders adds necessary headers.
func AddHeaders(c *fiber.Ctx) error {
	c.Response().Header.SetCanonical(altSVCHeader, altSVCValue)
	c.Response().Header.SetCanonical(hstsHeader, hstsValue)
	return c.Next()
}
