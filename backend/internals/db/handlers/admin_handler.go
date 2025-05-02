package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateAdminHandler)
	router.Get("/:id", middlewares.RequireAdmin, GetAdminByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateAdminHandler)
	router.Get("/", middlewares.RequireAdmin, ListAdminsHandler)
}

func CreateAdminHandler(c *fiber.Ctx) error {
	var admin operations.Admin
	if err := c.BodyParser(&admin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateAdmin(&admin)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"admin_id": id})
}

func GetAdminByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid admin ID")
	}

	admin, err := operations.GetAdminByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(admin)
}

func UpdateAdminHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid admin ID")
	}

	var updatedAdmin operations.Admin
	if err := c.BodyParser(&updatedAdmin); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateAdmin(id, &updatedAdmin); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update admin")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListAdminsHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	admins, err := operations.GetAdmins(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch admins")
	}

	return c.JSON(admins)
}
