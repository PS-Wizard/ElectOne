package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterAppealRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateAppealHandler)
	router.Get("/:id", middlewares.RequireAdmin, GetAppealByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateAppealHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteAppealHandler)
	router.Get("/", middlewares.RequireAdmin, ListAppealsHandler)
}

func CreateAppealHandler(c *fiber.Ctx) error {
	var appeal operations.Appeal
	if err := c.BodyParser(&appeal); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateAppeal(&appeal)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"appeal_id": id})
}

func GetAppealByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	appeal, err := operations.GetAppealByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(appeal)
}

func UpdateAppealHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	var updated operations.Appeal
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateAppeal(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update appeal")
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteAppealHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	if err := operations.DeleteAppeal(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete appeal")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListAppealsHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	appeals, err := operations.GetAppeals(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch appeals")
	}

	return c.JSON(appeals)
}
