package jwt

import (
	"rpiSite/models"

	JWTParser "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Protected is the function to make sure the user is logged in.
var Protected = func(c *fiber.Ctx) error {
	if _, ok := User(c); !ok {
		c.Status(fiber.StatusUnauthorized)
		return c.Render("login", fiber.Map{
			"Title": "Login is required",
			"Error": "You must log in to do this action.",
		})
	}
	return c.Next()
}

func mapClaim(c *fiber.Ctx) JWTParser.MapClaims {
	user, ok := c.Locals("user").(*JWTParser.Token)
	if !ok {
		return nil
	}
	claims := user.Claims.(JWTParser.MapClaims)

	return claims
}

// User function returns APIUser when a valid user is logged in otherwise empty.
func User(c *fiber.Ctx) (*models.APIUser, bool) {
	s := mapClaim(c)
	u := &models.APIUser{}

	if s == nil {
		return u, false
	}

	// Type assertion will convert interface{} to other types.
	u.Username = s["name"].(string)
	if s["email"] != nil {
		u.Email = s["email"].(string)
	}
	u.ID = uint(s["id"].(float64))
	u.Role = models.Role(s["role"].(float64))

	return u, true
}
