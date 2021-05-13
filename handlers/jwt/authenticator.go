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

	if name, ok := s["name"].(string); ok {
		u.Username = name
	} else {
		return u, false
	}

	// Email is optionial thus don't return false when it fails.
	if email, ok := s["email"].(string); ok {
		u.Email = email
	}

	if userID, ok := s["id"].(float64); ok {
		u.ID = uint(userID)
	} else {
		return u, false
	}

	if userRole, ok := s["role"].(float64); ok {
		u.Role = models.Role(userRole)
	} else {
		return u, false
	}

	return u, true
}
