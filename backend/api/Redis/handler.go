package redis

import "github.com/gofiber/fiber/v2"

func HandleVoteIncrement(ctx *fiber.Ctx) error {
	var voteRequest vote
	if err := ctx.BodyParser(&voteRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	err := incrementVote(voteRequest.ElectionID, voteRequest.CandidateID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to capture vote",
		})
	}
	return ctx.JSON(fiber.Map{
		"message": "Your Vote Has Been Captured",
	})
}
