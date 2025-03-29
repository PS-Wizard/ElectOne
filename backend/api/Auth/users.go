package auth

import (
	"database/sql"
	"time"

	"github.com/PS-Wizard/ElectOneAPI/api"
	users "github.com/PS-Wizard/ElectOneAPI/api/Users"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleUserLogin(ctx *fiber.Ctx) error {
	var req users.User

	// Parse the incoming request body into the User struct
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Fetch the stored password hash from the database
	var storedHash string
	var userID int
	query := "SELECT userID, password FROM users WHERE citizenID = ?"
	err := api.DB.QueryRow(query, req.CitizenID).Scan(&userID, &storedHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid credentials",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(req.Password)); err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"citizenID": req.CitizenID,
		"role":      "user",
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "user_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})

	return ctx.JSON(fiber.Map{"message": "Login successful"})
}
