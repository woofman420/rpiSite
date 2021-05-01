package common

import (
	"crypto/subtle"
	"log"
	"rpiSite/config"
	"rpiSite/utils"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func LoginGet(c *fiber.Ctx) error {
	if _, ok := c.Locals("user").(*jwt.Token); !ok {
		return c.Redirect("/index", fiber.StatusSeeOther)
	}

	return c.Render("login", fiber.Map{
		"Title": "Login",
	})
}

func LoginPost(c *fiber.Ctx) error {
	username, pwd := c.FormValue("username"), c.FormValue("password")
	remember := c.FormValue("remember") == "on"

	if subtle.ConstantTimeCompare(utils.S2b(username), utils.S2b(config.ADMIN_USER)) != 1 ||
		subtle.ConstantTimeCompare(utils.S2b(pwd), utils.S2b(config.ADMIN_PWD)) != 1 {
		log.Printf("Failed to match hash for user: %#+v\n", username)

		c.SendStatus(fiber.StatusInternalServerError)
		return c.Render("login", fiber.Map{
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
		SetClaim("name", username).
		SetExpiration(expiration).
		GetSignedString(nil)

	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return c.Render("err", fiber.Map{
			"Title": "Internal server error.",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     fiber.HeaderAuthorization,
		Value:    t,
		Path:     "/",
		Expires:  expiration,
		Secure:   config.IS_DEBUG == "false",
		HTTPOnly: true,
		SameSite: "strict",
	})

	return c.Redirect("/index", fiber.StatusSeeOther)
}
