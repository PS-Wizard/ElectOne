package handlers

import (
	"fmt"
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterUserRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateUserHandler)
	router.Get("/:id", GetUserByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateUserHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteUserHandler)
	router.Get("/", middlewares.RequireAdmin, ListUsersHandler)
	// router.Get("/lookup/", GetUserByCitizenAndVoterIDHandler)
}

func CreateUserHandler(c *fiber.Ctx) error {
	var user operations.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateUser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to create user %s", err))
	}

	return c.JSON(fiber.Map{"user_id": id})
}

func GetUserByIDHandler(c *fiber.Ctx) error {
	var id int
	var err error

	token := c.Locals("token").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
	}
	role := claims["role"]

	if role == "user" {
		user_id, ok := claims["user_id"].(float64)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid User Id")
		}
		id = int(user_id)
	} else {
		id, err = strconv.Atoi(c.Params("id"))
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
		}

	}

	user, err := operations.GetUserByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(user)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	var updated operations.User
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateUser(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	if err := operations.DeleteUser(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListUsersHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	users, err := operations.GetUsers(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch users")
	}

	return c.JSON(users)
}

func GetUserByCitizenAndVoterIDHandler(c *fiber.Ctx) error {
	citizenID := c.Query("citizen_id")
	voterID := c.Query("voter_card_id")

	fmt.Println(citizenID, voterID)
	if citizenID == "" || voterID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Missing citizen_id or voter_id")
	}

	user, err := operations.GetUserByCitizenAndVoterID(citizenID, voterID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}
