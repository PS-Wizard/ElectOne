package auth

import (
	"database/sql"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

var JWT_SECRET = []byte("supersecret123lol")

type Admin struct {
	AdminID  int    `json:"adminID"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleAdminLogin(ctx *fiber.Ctx) error {
	var req Admin

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var storedHash string
	var adminID int
	query := "SELECT adminID, password FROM admins WHERE email = ?"
	err := api.DB.QueryRow(query, req.Email).Scan(&adminID, &storedHash)
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"adminID": adminID,
		"email":   req.Email,
		"role":    "admin",
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	ctx.Cookie(&fiber.Cookie{
		Name:     "admin_token",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	})
	return ctx.JSON(fiber.Map{"message": "Login Successfull"})
}

func HandleCreateAdmin(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")

	if authHeader != "Bearer superadmin-token" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access, invalid token",
		})
	}

	var req Admin
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	query := "INSERT INTO admins (email, password) VALUES (?, ?)"
	_, err = api.DB.Exec(query, req.Email, string(hashedPassword))
	if err != nil {
		log.Println("Failed to insert admin:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create admin",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Admin created successfully",
	})
}
