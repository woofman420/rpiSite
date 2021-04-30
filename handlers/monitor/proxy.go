package monitor

import (
	"rpiSite/handlers/jwt"
	"rpiSite/models"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const addr = "http://localhost:19999"

var client = fasthttp.Client{
	NoDefaultUserAgentHeader: true,
	DisablePathNormalizing:   true,
}

func createAdress(c *fiber.Ctx) string {
	address := addr

	path := c.Path()[8:]
	address += path
	if path == "" {
		address += "/"
	}

	queryBytes := c.Context().QueryArgs().QueryString()
	if len(queryBytes) == 0 {
		return address
	}
	address += "?" + utils.B2s(queryBytes)
	return address
}

func ProxyMonitor(c *fiber.Ctx) error {
	u, ok := jwt.User(c)
	if !ok || u.Role != models.Admin {
		return c.Status(fiber.StatusUnauthorized).SendString("Not authorized!")
	}

	if c.Path() == "/monitor" {
		return c.Redirect("/monitor/", 301)
	}
	req := c.Request()
	res := c.Response()
	req.SetRequestURI(createAdress(c))

	host := utils.B2s(c.Context().Host())
	req.Header.Add("X-Forwarded-Host", host)
	req.Header.Add("X-Forwarded-Server", host)
	req.Header.Add("X-Forwarded-for", c.IP())
	if err := client.Do(req, res); err != nil {
		return err
	}
	return nil
}
