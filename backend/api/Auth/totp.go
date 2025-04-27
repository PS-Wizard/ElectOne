package auth

import (
	"fmt"
	"net/url"
	"time"

	"github.com/PS-Wizard/ElectOneAPI/api"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pquerna/otp/totp"
	"github.com/skip2/go-qrcode"
)

func generateTOTPSecret(citizenID string) (string, error) {

	issuer := "Electone."
	// Generate the TOTP secret
	otpKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: citizenID, //using the account name as the unique thing
	})
	if err != nil {
		return "", err
	}

	// Return the secret which can be used to generate the QR code
	return otpKey.Secret(), nil
}

func HandleGenerateQRCode(ctx *fiber.Ctx) error {
	// Get the OTP URL from the route parameter
	otpURL := ctx.Params("otpURL")

	// Decode the URL-encoded otpURL
	decodedURL, err := url.QueryUnescape(otpURL)
	if err != nil {
		fmt.Println("URL decode failed:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid OTP URL",
		})
	}

	// Log the decoded URL to see if it's empty or invalid
	fmt.Println("Decoded OTP URL:", decodedURL)

	// If the decoded URL is empty, return an error
	if decodedURL == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Empty OTP URL",
		})
	}

	// Try generating the QR code for the decoded URL
	qrCode, err := qrcode.Encode(decodedURL, qrcode.Medium, 256)
	if err != nil {
		fmt.Println("QR code generation failed:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate QR code",
		})
	}

	// Set the content type to image/png and return the QR code
	ctx.Set("Content-Type", "image/png")
	return ctx.Send(qrCode)
}

func generateTOTPQRCodeURL(citizenID, secret string) string {
	issuer := "Electone"
	otpURL := fmt.Sprintf(
		"otpauth://totp/%s:%s?secret=%s&issuer=%s",
		issuer,
		citizenID,
		secret,
		issuer,
	)
	encodedURL := url.QueryEscape(otpURL)

	return encodedURL
}

func ValidateTOTP(ctx *fiber.Ctx) error {
	otpCode := ctx.Query("otp") // Assuming OTP is passed as a query parameter
	citizenID := ctx.Query("citizenID")

	// Fetch the TOTP secret for the user
	var totpSecret string
	query := "SELECT totp_secret FROM users WHERE citizenID = ?"
	err := api.DB.QueryRow(query, citizenID).Scan(&totpSecret)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch TOTP secret",
		})
	}

	// Validate OTP using the stored secret
	valid := totp.Validate(otpCode, totpSecret)
	if !valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid OTP",
		})
	}

	// If OTP is valid, set the JWT cookie and return a success response
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"citizenID": citizenID,
		"role":      "user",
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	// Set the JWT cookie for the user
	ctx.Cookie(&fiber.Cookie{
		Name:     "user_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})

	// Return success message
	return ctx.JSON(fiber.Map{
		"message": "OTP validated successfully",
	})
}
