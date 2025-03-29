package users

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Get a user by ID
func HandleGetUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	if userID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, ID missing",
		})
	}
	user, err := getUser(userID)
	if err != nil {
		if err.Error() == "user not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(user)
}

// Get users paginated
func HandleGetUsersPaginated(ctx *fiber.Ctx) error {
	offset, err := ctx.ParamsInt("offset")
	if err != nil || offset < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset"})
	}
	users, err := getUsersPaginated(offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(users)
}

// Create a new user
func HandleCreateNewUser(ctx *fiber.Ctx) error {
	var u User
	if err := ctx.BodyParser(&u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := createNewUser(u)
	if err != nil {
		log.Printf("Error in HandleCreate: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

// Update user details
func HandleUpdateUserDetails(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	var u User
	if err := ctx.BodyParser(&u); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := updateUserDetails(u, userID)
	if err != nil {
		log.Printf("Error in HandleUpdate: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "User updated successfully"})
}

// Delete user
func HandleDeleteUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
    err := deleteUser(userID)
	if err != nil {
		log.Printf("Error in HandleDelete: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "User deleted successfully"})
}
