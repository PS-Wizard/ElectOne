package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterVoteRoutes(router fiber.Router) {
	router.Post("/", CreateVoteHandler)
	router.Get("/", middlewares.RequireAdmin, ListVotesHandler)
	router.Get("/voter/:voterCardID", GetVotesByVoterCardHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteVoteHandler)
	router.Get("/history/:electionID", GetVotesByElectionIDHandler)
}

func CreateVoteHandler(c *fiber.Ctx) error {
	var vote operations.Vote
	if err := c.BodyParser(&vote); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID")
	}
	vote.VoterCardID = voterCardID
	id, err := operations.CreateVote(&vote)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"vote_id": id})
}

func ListVotesHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	votes, err := operations.GetVotes(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch votes")
	}

	return c.JSON(votes)
}

func GetVotesByVoterCardHandler(c *fiber.Ctx) error {
	voterCardID := c.Params("voterCardID")

	votes, err := operations.GetVotesByVoterCard(voterCardID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(votes)
}

func DeleteVoteHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid vote ID")
	}

	if err := operations.DeleteVote(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func GetVotesByElectionIDHandler(c *fiber.Ctx) error {
	electionID, err := strconv.Atoi(c.Params("electionID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid election ID")
	}

	limit, _ := strconv.Atoi(c.Query("limit", "100"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	votes, err := operations.GetVotesByElectionID(electionID, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch votes")
	}

	return c.JSON(votes)
}
