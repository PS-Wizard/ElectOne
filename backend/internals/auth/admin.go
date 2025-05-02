package auth

import (
	"time"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	// "github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// TOTPCode string `json:"totp_code"`
}

func AdminLogin(c *fiber.Ctx) error {
	var input AdminLoginInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	admin, err := operations.GetAdminByEmail(input.Email)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch admin")
	}
	if admin == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// TOTP check
	// if admin.TotpSecret != "" {
	// 	valid := totp.Validate(input.TOTPCode, admin.TotpSecret)
	// 	if !valid {
	// 		return fiber.NewError(fiber.StatusUnauthorized, "Invalid TOTP code")
	// 	}
	// }

	// Token generation
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
		"token": tokenStr,
	})
}
