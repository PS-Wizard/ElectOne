package routes

import (
	"log"

	auth "github.com/PS-Wizard/ElectOneAPI/api/Auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TokenValidationAdmin(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("admin_token")
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization cookie is missing",
		})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
		}
		return auth.JWT_SECRET, nil
	})

	if err != nil || !token.Valid {
		log.Println("Invalid or expired token")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role := claims["role"].(string)
		if role != "admin" {
			log.Println("Unauthorized role")
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized role",
			})
		}
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	return ctx.Next()
}

func TokenValidationUser(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token != "Bearer usersecrettoken" {
		log.Println("Invalid User Token")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access, invalid user token",
		})
	}

	return ctx.Next()
}
