package monitor

import (
	"rpiSite/handlers/jwt"
	"rpiSite/models"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

var (
	addr = []byte("http://localhost:19999")
)

var client = fasthttp.Client{
	NoDefaultUserAgentHeader:      true,
	DisablePathNormalizing:        true,
	DisableHeaderNamesNormalizing: true,
}

func createAdrress(c *fiber.Ctx) (address []byte) {
	path := c.Request().URI().Path()[8:]
	address = append(addr, path...)
	if len(path) == 0 {
		// Byte 47 is `/`
		address = append(address, 47)
	}

	if c.Context().QueryArgs().Len() == 0 {
		return address
	}
	// Byte 63 = `?`
	address = append(address, 63)
	address = append(address, c.Context().QueryArgs().QueryString()...)
	return
}

// ProxyMonitor is the handler for `/monitor`
// It does the hardwork of reverse-proxing the local
// netdata installation.
func ProxyMonitor(c *fiber.Ctx) error {
	u, ok := jwt.User(c)
	if !ok || u.Role != models.Admin {
		return c.Status(fiber.StatusUnauthorized).
			SendString("Not authorized!")
	}

	// String equality is faster than bytes equality so we don't optimizate this.
	if c.Path() == "/monitor" {
		return c.Redirect("/monitor/", 301)
	}
	req := c.Request()
	req.SetRequestURIBytes(createAdrress(c))

	return client.Do(req, c.Response())
}
