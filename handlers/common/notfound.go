package common

import "github.com/gofiber/fiber/v2"

func NotFound(c *fiber.Ctx) error {
	return c.Render("err", fiber.Map{
		"Title": "Page not found",
		"Error": "404 - Your request wasn't found (゜ロ゜)",
	})
}
