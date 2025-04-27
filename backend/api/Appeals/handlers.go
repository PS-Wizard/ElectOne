package appeals

import "github.com/gofiber/fiber/v2"

func HandleGetRegistrationAppeals(ctx *fiber.Ctx) error {
	appeals, err := getRegistrationAppeals()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(appeals)
}

func HandleAppealApprove(ctx *fiber.Ctx) error {
	appealID := ctx.Params("id")
	err := approveRegistrationAppeal(appealID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Appeal Successfully Approved"})
}
