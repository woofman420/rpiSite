package jwt

import (
	"rpiSite/models"

	JWTParser "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Protected is the function to make sure the user is logged in.
var Protected = func(c *fiber.Ctx) error {
	if _, ok := User(c); !ok {
		return c.Status(fiber.StatusUnauthorized).
			Render("login", fiber.Map{
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
	claims, ok := user.Claims.(JWTParser.MapClaims)
	if !ok {
		return nil
	}

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

	name, ok := s["name"].(string)
	if !ok {
		return u, false
	}
	u.Username = name

	// Email is optionial thus don't return false when it fails.
	if email, ok := s["email"].(string); ok {
		u.Email = email
	}
	userID, ok := s["id"].(float64)
	if !ok {
		return u, false
	}
	u.ID = uint(userID)

	userRole, ok := s["role"].(float64)
	if !ok {
		return u, false
	}
	u.Role = models.Role(userRole)

	return u, true
}
