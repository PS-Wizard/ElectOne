package routes

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/PS-Wizard/ElectOneAPI/api"
	auth "github.com/PS-Wizard/ElectOneAPI/api/Auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
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

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Extract useful claims (userID, citizenID, role, etc.)
	userID := int(claims["userID"].(float64)) // Assuming userID is stored as float64 in the JWT
	citizenID := claims["citizenID"].(string)
	role := claims["role"].(string)

	// Store the extracted data in context for use in subsequent handlers
	ctx.Locals("userID", userID)
	ctx.Locals("citizenID", citizenID)
	ctx.Locals("role", role)

	// Check for roles, if needed
	if role != "user" {
		log.Println("Unauthorized role")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized role",
		})
	}

	// Continue to the next handler
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

func TOTPValidation(ctx *fiber.Ctx) error {
	// Get the OTP from the request body
	var req struct {
		OTP string `json:"otp"` // OTP provided by the user
	}

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Extract citizenID from the JWT token
	citizenID := ctx.Locals("citizenID").(string) // This comes from the TokenValidationUser middleware

	// Fetch the stored TOTP secret from the database
	var storedSecret string
	query := "SELECT totp_secret FROM users WHERE citizenID = ?"
	err := api.DB.QueryRow(query, citizenID).Scan(&storedSecret)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found or TOTP not enabled",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	// Validate the OTP using the stored secret
	valid := totp.Validate(req.OTP, storedSecret)
	if !valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid OTP",
		})
	}

	// OTP is valid, allow the user to proceed
	return ctx.JSON(fiber.Map{
		"message": "OTP validated successfully",
	})
}
