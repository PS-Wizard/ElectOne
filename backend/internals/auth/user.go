package auth

import (
	"time"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginRequest struct {
	CitizenshipID string `json:"citizenship_id"`
	VoterCardID   string `json:"voter_card_id"`
	Password      string `json:"password"`
	// TOTP          string `json:"totp"`
}

func UserLogin(c *fiber.Ctx) error {
	var req UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	user, err := operations.GetUserByCitizenAndVoterID(req.CitizenshipID, req.VoterCardID)
	if err != nil || user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	// if user.TotpSecret != "" && !totp.Validate(req.TOTP, user.TotpSecret) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid TOTP"})
	// }

	claims := jwt.MapClaims{
		"user_id":     user.UserID,
		"citizenship": user.CitizenshipID,
		"voter_card":  user.VoterCardID,
		"role":        "user",
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(JWTSECRET))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Token creation failed"})
	}

	return c.JSON(fiber.Map{
		"token": signedToken,
	})
}
