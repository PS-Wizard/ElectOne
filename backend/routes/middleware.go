package routes

import (
	"fmt"
	"log"

	auth "github.com/PS-Wizard/ElectOneAPI/api/Auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func TokenValidationAdmin(ctx *fiber.Ctx) error {
	tokenString := ctx.Cookies("admin_token")
	fmt.Println("got admin_token: ", tokenString)
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Admin Authorization cookie is missing",
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
	// Get the token from the user_token cookie
	tokenString := ctx.Cookies("user_token")
	fmt.Println("Got user_token: ", tokenString)
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User Authorization cookie is missing",
		})
	}

	// Parse the token using the same secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
		}
		return auth.JWT_SECRET, nil
	})

	// Check if there's an error or if the token is invalid
	if err != nil || !token.Valid {
		log.Println("Invalid or expired token")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Verify the token claims (you can also check the role if needed)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role := claims["role"].(string)
		if role != "user" {
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

func TokenValidationBoth(ctx *fiber.Ctx) error {
	adminToken := ctx.Cookies("admin_token")
	if adminToken != "" {
		return TokenValidationAdmin(ctx)
	}

	userToken := ctx.Cookies("user_token")
	if userToken != "" {
		// If a user token exists, validate it
		return TokenValidationUser(ctx)
	}

	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "Authorization cookie is missing",
	})
}
