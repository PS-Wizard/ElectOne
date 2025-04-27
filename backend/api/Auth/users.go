package auth

import (
	"database/sql"

	"github.com/PS-Wizard/ElectOneAPI/api"
	users "github.com/PS-Wizard/ElectOneAPI/api/Users"
	"github.com/gofiber/fiber/v2"
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

	// Fetch the stored password hash and TOTP secret from the database
	var storedHash string
	var userID int
	var totpSecret sql.NullString // Use sql.NullString to handle NULL values in the DB
	query := "SELECT userID, password, totp_secret FROM users WHERE citizenID = ?"
	err := api.DB.QueryRow(query, req.CitizenID).Scan(&userID, &storedHash, &totpSecret)
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

	// If TOTP is not set (NullString check), generate the secret and return it along with the login response
	if !totpSecret.Valid {
		// Generate the TOTP secret
		secret, err := generateTOTPSecret(req.CitizenID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to generate TOTP secret",
			})
		}

		// Store the TOTP secret in the database
		_, err = api.DB.Exec("UPDATE users SET totp_secret = ? WHERE citizenID = ?", secret, req.CitizenID)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update TOTP secret in database",
			})
		}

		// Return the TOTP setup QR code URL or other details
		otpURL := generateTOTPQRCodeURL(req.CitizenID, secret)

		return ctx.JSON(fiber.Map{
			"message":     "Login successful, please set up your 2FA",
			"totp_secret": otpURL, // send the QR code URL or other relevant data
		})
	}

	// If TOTP is already set, return a message indicating that OTP is required for verification
	return ctx.JSON(fiber.Map{
		"message": "Credentials matched, please enter your OTP to verify",
	})
}
