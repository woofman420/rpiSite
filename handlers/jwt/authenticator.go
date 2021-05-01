package jwt

import (
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var Protected = func(c *fiber.Ctx) error {
	if _, ok := c.Locals("user").(*jwt.Token); !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.Render("login", fiber.Map{
			"Title": "Login is required",
			"Error": "You must log in to do this action.",
		})
	}
	return c.Next()
}
