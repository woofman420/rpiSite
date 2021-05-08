package common

import "github.com/gofiber/fiber/v2"

// Index is our handler for the `/`.
func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}
