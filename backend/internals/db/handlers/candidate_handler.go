package handlers

import (
	"fmt"
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/PS-Wizard/Electone/utils"
	"github.com/gofiber/fiber/v2"
)

func RegisterCandidateRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateCandidateHandler)
	router.Get("/:id", GetCandidatesByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateCandidateHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteCandidateHandler)
	router.Get("/", ListCandidatesHandler)
}

func CreateCandidateHandler(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid form data")
	}

	files := form.File["candidate_photo"]
	if len(files) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "No photo file provided")
	}

	file := files[0]
	uniqueName := utils.GenerateUniqueFileName(file.Filename)

	// Upload to S3
	if err := utils.UploadToS3(file, uniqueName); err != nil {
		fmt.Println("S3 Upload Error:", err)
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload file to S3")
	}

	// Parse election_id
	electionID, err := strconv.Atoi(c.FormValue("election_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Election ID")
	}

	candidate := operations.Candidate{
		CitizenID:   c.FormValue("citizenship_id"),
		ElectionID:  electionID,
		ProfilePath: utils.BUCKETURL + uniqueName,
		Bio:         c.FormValue("candidate_bio"),
		Post:        c.FormValue("candidate_post"),
		Party:       c.FormValue("candidate_party"),
		Name:        c.FormValue("candidate_name"),
	}

	id, err := operations.CreateCandidate(&candidate)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to Create Candidate: %s", err.Error()),
		})
	}

	return c.JSON(fiber.Map{"candidate_id": id})
}


func GetCandidatesByIDHandler(c *fiber.Ctx) error {
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
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Candidate ID")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid form data")
	}

	var profilePath string
	candidatePhoto := form.File["candidate_photo"]

	if len(candidatePhoto) > 0 {
		uniqueName := utils.GenerateUniqueFileName(candidatePhoto[0].Filename)
		if err := utils.UploadToS3(candidatePhoto[0], uniqueName); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to upload photo")
		}
		profilePath = utils.BUCKETURL + uniqueName
	} else {
		existingCandidate, err := operations.GetCandidateByID(id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch existing candidate data")
		}
		profilePath = existingCandidate.ProfilePath
	}

	electionIDStr := c.FormValue("election_id")
	electionID, err := strconv.Atoi(electionIDStr)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid Election ID")
	}

	updatedCandidate := operations.Candidate{
		CitizenID:   c.FormValue("citizenship_id"),
		ElectionID:  electionID,
		ProfilePath: profilePath,
		Bio:         c.FormValue("candidate_bio"),
		Post:        c.FormValue("candidate_post"),
		Party:       c.FormValue("candidate_party"),
		Name:        c.FormValue("candidate_name"),
	}

	if err := operations.UpdateCandidate(id, &updatedCandidate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to update Candidate: %s", err.Error()),
		})
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
