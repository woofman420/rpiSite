package jwt

import (
	"strings"

	"rpiSite/config"
	"rpiSite/utils"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	signingKey    = []byte(config.JWTSigningKey)
	signingMethod = "HS512"
)

func keyFunction(t *jwt.Token) (interface{}, error) {
	if t.Method.Alg() != signingMethod {
		return nil, utils.ErrorUnexpectedJWTAlgoirthm(t.Method.Alg())
	}
	return signingKey, nil
}

// New function to create a new JWT middleware.
func New() fiber.Handler {
	extractors := []func(c *fiber.Ctx) (string, bool){
		jwtFromCookie(fiber.HeaderAuthorization),
		jwtFromHeader(fiber.HeaderAuthorization),
	}

	return func(c *fiber.Ctx) error {
		var auth string
		var ok bool

		for _, extractor := range extractors {
			auth, ok = extractor(c)
			if auth != "" && ok {
				break
			}
		}

		if !ok {
			return c.Next()
		}

		token, err := jwt.Parse(auth, keyFunction)

		if err == nil && token.Valid {
			// Store user information from token into context.
			c.Locals("user", token)
			return c.Next()
		}
		return c.Next()
	}
}

func jwtFromHeader(header string) func(c *fiber.Ctx) (string, bool) {
	return func(c *fiber.Ctx) (string, bool) {
		auth := c.Get(header)
		l := len("Bearer")
		if len(auth) > l+1 && strings.EqualFold(auth[:l], "Bearer") {
			return auth[l+1:], true
		}
		return "", false
	}
}

func jwtFromCookie(name string) func(c *fiber.Ctx) (string, bool) {
	return func(c *fiber.Ctx) (string, bool) {
		token := c.Cookies(name)
		if token == "" {
			return "", false
		}
		return token, true
	}
}
