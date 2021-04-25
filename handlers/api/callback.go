package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func CallbackGet(c *fiber.Ctx) error {
	fmt.Println("aa")
	return c.Render("callback_helper", fiber.Map{
		"Callback": c.Request().URI().QueryArgs().String(),
	})
}
