package auth

import (
	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
	"time"

	// "github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	TOTPCode string `json:"totp_code"`
}

func AdminLogin(c *fiber.Ctx) error {
	var input AdminLoginInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	admin, err := operations.GetAdminByEmail(input.Email)
	if err != nil || admin == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Password check
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// TOTP setup
	if admin.TotpSecret == "" {
		// Generate new secret
		key, err := totp.Generate(totp.GenerateOpts{
			Issuer:      "Electone",
			AccountName: admin.Email,
		})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to generate TOTP")
		}

		// Save secret to DB
		if err := operations.UpdateAdminTotpSecret(admin.AdminID, key.Secret()); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not save TOTP secret")
		}

		// Return QR code URI (user scans with authenticator app)
		return c.JSON(fiber.Map{
			"message":    "Scan this QR code to set up TOTP",
			"qr_url":     key.URL(),
			"setup_done": false,
		})
	}

	// TOTP validation
	if !totp.Validate(input.TOTPCode, admin.TotpSecret) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid TOTP code")
	}

	// Create JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": admin.AdminID,
		"email":   admin.Email,
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(JWTSECRET))
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not sign token")
	}

	return c.JSON(fiber.Map{
		"token":      tokenStr,
		"setup_done": true,
	})
}
