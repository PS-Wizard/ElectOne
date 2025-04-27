package citizens

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleSearch(ctx *fiber.Ctx) error {
	citizenID := ctx.Params("id")
	if citizenID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen ID is required",
		})
	}

	citizen, err := getCitizenById(citizenID)
	if err != nil {
		log.Println("Error Fetching Citizen: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen with that ID not found",
		})
	}
	return ctx.JSON(citizen)
}

func HandleGet(ctx *fiber.Ctx) error {
	offset, err := ctx.ParamsInt("offset")
	if err != nil || offset < 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid offset"})
	}
	citizens, err := getCitizensPaginated(offset)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(citizens)
}

func HandleCreate(ctx *fiber.Ctx) error {
	var citizen Citizen
	if err := ctx.BodyParser(&citizen); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := CreateNewCitizen(citizen)
	if err != nil {
		log.Printf("Error in HandleCreate: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create citizen",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Citizen created successfully",
	})
}

func HandleUpdate(ctx *fiber.Ctx) error {
	citizenID := ctx.Params("id")
	if citizenID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen ID is required",
		})
	}

	var citizen Citizen
	if err := ctx.BodyParser(&citizen); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := updateCitizenDetails(citizen, citizenID)
	if err != nil {
		log.Printf("Error in updateCitizenDetails: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update citizen",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Citizen updated successfully",
	})
}

func HandleDelete(ctx *fiber.Ctx) error {
	citizenID := ctx.Params("id")
	if citizenID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen ID is required",
		})
	}
	err := deleteCitizen(citizenID)
	if err != nil {
		log.Printf("Error in deleteCitizen: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete citizen",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Citizen deleted successfully",
	})
}

