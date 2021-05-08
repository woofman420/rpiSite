package common

import "github.com/gofiber/fiber/v2"

// NotFound is our handler when no other handler rendered or returned something.
func NotFound(c *fiber.Ctx) error {
	c.Status(404)
	return c.Render("err", fiber.Map{
		"Title": "Page not found",
		"Error": "404 - Your request wasn't found (゜ロ゜)",
	})
}
