package elections

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleGetElection(ctx *fiber.Ctx) error {
	electionID := ctx.Params("id")
	if electionID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, ID missing",
		})
	}
	election, err := getElectionDetail(electionID)
	if err != nil {
		if err.Error() == "election not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "election not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(election)
}
func HandleGetElectionsPaginated(ctx *fiber.Ctx) error {
	offset, err := ctx.ParamsInt("offset")
	if offset < 0 || err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid Request, ID missing",
		})
	}
	elections, err := getElectionsPaginated(offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(elections)
}

func HandleCreateNewElection(ctx *fiber.Ctx) error {
	var details Election
	if err := ctx.BodyParser(&details); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := createElection(details)
	if err != nil {
		log.Printf("Error in Election Creation: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Election successfully Created"})
}

func HandleUpdateElectionDetails(ctx *fiber.Ctx) error {
	electionId := ctx.Params("id")
	if electionId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var details Election
	if err := ctx.BodyParser(&details); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := updateElectionDetail(electionId, details)
	if err != nil {
		log.Printf("Error in Election Creation: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Election Details Updated"})
}

func HandleDeleteElection(ctx *fiber.Ctx) error {
	electionId := ctx.Params("id")
	if electionId == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	err := deleteElection(electionId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Election Deleted"})
}
