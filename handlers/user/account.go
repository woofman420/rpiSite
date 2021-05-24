package user

import (
	"rpiSite/handlers/jwt"

	"github.com/gofiber/fiber/v2"
)

// AccountGet is the handler for `GET /account`
func AccountGet(c *fiber.Ctx) error {
	user, _ := jwt.User(c)

	return c.Render("account", fiber.Map{
		"User": user,
	})
}
