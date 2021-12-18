package common

import "github.com/gofiber/fiber/v2"

// PortfolioGet is our handler for `GET /2021/portfolio`.
func PortfolioGet(c *fiber.Ctx) error {
	return c.Render("portfolio", fiber.Map{})
}
