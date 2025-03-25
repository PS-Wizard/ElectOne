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

	citizen, err := GetCitizenByID(citizenID, api.DB)
	if err != nil {
		log.Println("Error Fetching Citizen: ", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Citizen with that ID not found",
		})
	}
	return ctx.JSON(citizen)
}

func HandleGet(ctx *fiber.Ctx) error {

	return nil
}

func HandleCreate(ctx *fiber.Ctx) error {

	return nil
}

func HandleUpdate(ctx *fiber.Ctx) error {

	return nil
}

func HandleDelete(ctx *fiber.Ctx) error {

	return nil
}
