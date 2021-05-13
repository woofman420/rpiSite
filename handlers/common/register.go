package common

import (
	"crypto/subtle"
	"log"
	"rpiSite/config"
	"rpiSite/database"
	"rpiSite/handlers/jwt"
	"rpiSite/models"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
)

// RegisterGet is our handler for `GET /register`.
func RegisterGet(c *fiber.Ctx) error {
	if u, ok := jwt.User(c); ok {
		log.Printf("User %d has set session, redirecting.", u.ID)
		return c.Redirect("/account", fiber.StatusSeeOther)
	}

	return c.Render("register", fiber.Map{
		"Title": "Register",
	})
}

// RegisterPost is our handler for `POST /register`.
func RegisterPost(c *fiber.Ctx) error {
	secretCode := c.FormValue("secret_code")

	if subtle.ConstantTimeCompare(utils.UnsafeByteConversion(secretCode),
		utils.UnsafeByteConversion(config.SecretCode)) != 1 {
		return c.Render("err", fiber.Map{
			"Error": "Woopsie wrong secret code",
		})
	}

	u := models.User{
		Username: c.FormValue("username"),
		Password: utils.GenerateHashedPassword(c.FormValue("password")),
		Email:    c.FormValue("email"),
	}

	regErr := database.DB.Create(&u)

	if regErr.Error != nil {
		log.Printf("Failed to register %s, error: %s", u.Email, regErr.Error)

		return c.Status(fiber.StatusInternalServerError).
			Render("err", fiber.Map{
				"Title": "Register failed",
				"Error": "Internal server error.",
			})
	}

	return c.Redirect("/login", fiber.StatusSeeOther)
}
