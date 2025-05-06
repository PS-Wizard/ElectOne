package handlers

import (
	"fmt"
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/PS-Wizard/Electone/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterAppealRoutes(router fiber.Router) {
	router.Post("/", CreateAppealHandler)
	router.Get("/:id", middlewares.RequireAdmin, GetAppealByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateAppealHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteAppealHandler)
	router.Get("/", middlewares.RequireAdmin, ListAppealsHandler)

	router.Post("/:id/approve", middlewares.RequireAdmin, ApproveAppealHandler)
	router.Post("/:id/reject", middlewares.RequireAdmin, RejectAppealHandler)
}

func CreateAppealHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid form data")
	}

	files := form.File["photos"]
	if len(files) != 4 {
		return fiber.NewError(fiber.StatusBadRequest, "Exactly 4 photos must be uploaded")
	}

	var photoPaths []string

	for _, photo := range files {
		uniqueName := utils.GenerateUniqueFileName(photo.Filename)
		path := fmt.Sprintf("./uploads/photos/%s", uniqueName)

		if err := c.SaveFile(photo, path); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed To Save File")
		}

		photoPaths = append(photoPaths, "/uploads/photos/"+uniqueName)
	}

	appeal := operations.Appeal{
		CitizenshipID: c.FormValue("citizenship_id"),
		VoterCardID:   c.FormValue("voter_card_id"),
		Password:      c.FormValue("password"),
		Photos:        fmt.Sprintf("%s", utils.Join(photoPaths, ",")),
		Status:        "pending",
	}

	id, err := operations.CreateAppeal(&appeal)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"appeal_id": id})
}

func GetAppealByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	appeal, err := operations.GetAppealByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Appeal not found")
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
	limit, err1 := strconv.Atoi(c.Query("limit", "10"))
	offset, err2 := strconv.Atoi(c.Query("offset", "0"))

	if err1 != nil || err2 != nil || limit < 0 || offset < 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid pagination params")
	}

	appeals, err := operations.GetAppeals(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch appeals")
	}

	return c.JSON(appeals)
}

func ApproveAppealHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	if err := operations.ApproveAppeal(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to approve appeal: %v", err))
	}

	return c.JSON(fiber.Map{
		"message": "Appeal approved and user registered successfully",
	})
}

func RejectAppealHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid appeal ID")
	}

	if err := operations.DeleteAppeal(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, fmt.Sprintf("Failed to reject appeal: %v", err))
	}

	return c.JSON(fiber.Map{
		"message": "Appeal rejected and removed",
	})
}
