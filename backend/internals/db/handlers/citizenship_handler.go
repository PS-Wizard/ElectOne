package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterCitizenshipRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateCitizenshipHandler)
	router.Get("/:id", GetCitizenshipByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateCitizenshipHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteCitizenshipHandler)
	router.Get("/", middlewares.RequireAdmin, ListCitizenshipsHandler)
}

func CreateCitizenshipHandler(c *fiber.Ctx) error {
	var citizen operations.Citizenship
	if err := c.BodyParser(&citizen); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateCitizenship(&citizen)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"citizen_id": id})
}

func GetCitizenshipByIDHandler(c *fiber.Ctx) error {
	var id string
	token := c.Locals("token").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
	}

	role := claims["role"]

	if role == "user" {
		id = claims["citizenship"].(string)
	} else {
		id = c.Params("id")
	}

	citizen, err := operations.GetCitizenshipByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(citizen)
}

func UpdateCitizenshipHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var updated operations.Citizenship
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateCitizenship(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update citizenship")
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteCitizenshipHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := operations.DeleteCitizenship(db.DB, id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete citizenship")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListCitizenshipsHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	citizens, err := operations.GetCitizenships(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch records")
	}

	return c.JSON(citizens)
}
