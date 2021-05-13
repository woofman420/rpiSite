package api

import (
	"github.com/gofiber/fiber/v2"
)

// CallbackGet is the handler for `GET /callback_helper`.
func CallbackGet(c *fiber.Ctx) error {
	return c.Render("callback_helper", fiber.Map{
		"Callback": c.Request().URI().QueryArgs().String(),
	})
}
