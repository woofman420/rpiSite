package common

import (
	"log"
	"time"

	"rpiSite/config"
	"rpiSite/database"
	"rpiSite/handlers/jwt"
	"rpiSite/models"
	"rpiSite/utils"

	"github.com/gofiber/fiber/v2"
)

// LoginGet is our handler for `GET /login`.
func LoginGet(c *fiber.Ctx) error {
	if u, ok := jwt.User(c); ok {
		log.Printf("User %d has set session, redirecting.", u.ID)
		return c.Redirect("/account", fiber.StatusSeeOther)
	}

	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

// LoginPost is our handler for `GET /post`.
func LoginPost(c *fiber.Ctx) error {
	form := models.User{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	remember := c.FormValue("remember") == "on"

	user, err := models.FindUserByEmail(database.DB, form.Email)
	if err != nil {
		log.Printf("Failed to find %s, error: %s", form.Email, err)

		return c.Status(fiber.StatusUnauthorized).
			Render("login", fiber.Map{
				"Title": "Login failed",
				"Error": "Invalid credentials.",
			})
	}

	if !utils.CompareHashedPassword(user.Password, form.Password) {
		log.Printf("Failed to match hash for user: %#+v\n", user.Email)

		return c.Status(fiber.StatusInternalServerError).
			Render("login", fiber.Map{
				"Title": "Login failed",
				"Error": "Invalid credentials.",
			})
	}

	var expiration time.Time
	if remember {
		// 2 weeks
		expiration = time.Now().Add(time.Hour * 24 * 14)
	}
	t, err := utils.NewJWTToken().
		SetClaim("id", user.ID).
		SetClaim("name", user.Username).
		SetClaim("email", user.Email).
		SetClaim("role", user.Role).
		SetExpiration(expiration).
		GetSignedString(nil)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			Render("err", fiber.Map{
				"Title": "Internal server error.",
			})
	}

	c.Cookie(&fiber.Cookie{
		Name:     fiber.HeaderAuthorization,
		Value:    t,
		Path:     "/",
		Expires:  expiration,
		Secure:   !config.IsDebug,
		HTTPOnly: true,
		SameSite: "strict",
	})

	return c.Redirect("/account", fiber.StatusSeeOther)
}
