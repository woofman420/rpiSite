package common

import "github.com/gofiber/fiber/v2"

// EindgesprekGet is our handler for `GET /2021/eindgesprek`.
func EindgesprekGet(c *fiber.Ctx) error {
	return c.Render("eindgesprek", fiber.Map{})
}
