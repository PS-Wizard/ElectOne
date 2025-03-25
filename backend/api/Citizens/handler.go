package citizens

import (
	"log"

	"github.com/PS-Wizard/ElectOneAPI/api"
	"github.com/gofiber/fiber/v2"
)

func HandleSearch(ctx *fiber.Ctx) error {
	citizenID := ctx.Params("id")
	if citizenID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen ID is required",
		})
	}

	citizen, err := getCitizenById(citizenID, api.DB)
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
	var citizen citizen
	if err := ctx.BodyParser(&citizen); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	query := "INSERT INTO citizens (citizenID, fullName, dateOfBirth, placeOfResidence) VALUES (?, ?, ?, ?)"
	_, err := api.DB.Exec(query, citizen.CitizenID, citizen.Name, citizen.DOB, citizen.POR)
	if err != nil {
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

	var citizen citizen
	if err := ctx.BodyParser(&citizen); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	query := "UPDATE citizens SET fullName = ?, dateOfBirth = ?, placeOfResidence = ? WHERE citizenID = ?"
	_, err := api.DB.Exec(query, citizen.Name, citizen.DOB, citizen.POR, citizenID)
	if err != nil {
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

	query := "DELETE FROM citizens WHERE citizenID = ?"
	_, err := api.DB.Exec(query, citizenID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete citizen",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Citizen deleted successfully",
	})
}
