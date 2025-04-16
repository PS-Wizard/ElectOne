package candidates

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleGetCandidate(ctx *fiber.Ctx) error {
	candidateID := ctx.Params("id")
	if candidateID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, ID missing",
		})
	}

	candidate, err := getCandidateDetail(candidateID)
	if err != nil {
		if err.Error() == "candidate not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Candidate not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(candidate)
}

func HandleGetCandidatesPaginated(ctx *fiber.Ctx) error {
	offset, err := ctx.ParamsInt("offset")
	if offset < 0 || err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, Offset missing or invalid",
		})
	}

	candidates, err := getCandidatesPaginated(offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(candidates)
}

func HandleCreateCandidate(ctx *fiber.Ctx) error {
	var candidate Candidate
	if err := ctx.BodyParser(&candidate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := createCandidate(candidate)
	if err != nil {
		log.Printf("Error in Candidate Creation: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Candidate successfully created"})
}

func HandleUpdateCandidate(ctx *fiber.Ctx) error {
	candidateID := ctx.Params("id")
	if candidateID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var candidate Candidate
	if err := ctx.BodyParser(&candidate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := updateCandidate(candidateID, candidate)
	if err != nil {
		log.Printf("Error in Candidate Update: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Candidate updated successfully"})
}

func HandleDeleteCandidate(ctx *fiber.Ctx) error {
	candidateID := ctx.Params("id")
	if candidateID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	err := deleteCandidate(candidateID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Candidate deleted"})
}
