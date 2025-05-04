package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterCandidateRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateCandidateHandler)
	router.Get("/:id", GetCandidateByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateCandidateHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteCandidateHandler)
	router.Get("/", ListCandidatesHandler)
}

func CreateCandidateHandler(c *fiber.Ctx) error {
	var candidate operations.Candidate
	if err := c.BodyParser(&candidate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateCandidate(&candidate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"candidate_id": id})
}

func GetCandidateByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid candidate ID")
	}

	candidate, err := operations.GetCandidateByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(candidate)
}

func UpdateCandidateHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid candidate ID")
	}

	var updated operations.Candidate
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateCandidate(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteCandidateHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid candidate ID")
	}

	if err := operations.DeleteCandidate(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete candidate")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListCandidatesHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	candidates, err := operations.GetCandidates(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch candidates")
	}

	return c.JSON(candidates)
}
