package api

import (
	"github.com/gofiber/fiber/v2"
)

// CallbackGet is the handler for `GET /callback_helper/:type?`
func CallbackGet(c *fiber.Ctx) error {
	typeCallback := c.Params("type")

	if typeCallback == "usw" {
		return c.Render("callback_helper", fiber.Map{
			"Callback": c.Request().URI().QueryArgs().String(),
			"USw":      true,
		})
	}
	return c.Render("callback_helper", fiber.Map{
		"Callback": c.Request().URI().QueryArgs().String(),
	})
}
