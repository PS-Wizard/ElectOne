package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func TokenValidationAdmin(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")

	if token != "Bearer adminsecrettoken" {
		log.Println("Invalid Token")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized access, invalid token",
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
