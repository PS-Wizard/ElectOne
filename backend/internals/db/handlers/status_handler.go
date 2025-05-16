package handlers

import (
	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
)

func RegisterStatusRoutes(router fiber.Router) {
	router.Post("/", HandleAppealGetStatus)

}

func HandleAppealGetStatus(ctx *fiber.Ctx) error {
	var request operations.AppealStatusReq

	if err := ctx.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Request Body")
	}

	appeal, err := operations.GetAppealStatus(&request)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error()) 
	}

	if appeal == nil {
		return fiber.NewError(fiber.StatusNotFound, "No Record Found")
	}

	return ctx.JSON(appeal)
}
