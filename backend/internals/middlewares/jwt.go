package middlewares

import (
	"fmt"

	"github.com/PS-Wizard/Electone/internals/auth"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(auth.JWTSECRET)},
		ContextKey:   "token",
		ErrorHandler: jwtError,
	})
}

func jwtError(ctx *fiber.Ctx, err error) error {
	fmt.Println("[JWT ERROR]", err) // 👈 add this
	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}

func RequireAdmin(c *fiber.Ctx) error {
	token := c.Locals("token")
	if token == nil {
		// Handle missing token
		return jwtError(c, fmt.Errorf("missing or malformed JWT"))
	}

	parsedToken, ok := token.(*jwt.Token)
	if !ok {
		return jwtError(c, fmt.Errorf("invalid token format"))
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return jwtError(c, fmt.Errorf("invalid claims"))
	}

	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Nuh Uh Admins Only",
		})
	}

	return c.Next()
}

func RequireUser(c *fiber.Ctx) error {
	claims := c.Locals("token").(jwt.MapClaims)
	if claims["role"] != "user" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Nuh Uh Users Only"})
	}
	return c.Next()
}
