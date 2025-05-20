package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/ws"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterRedisVotingRoutes(app fiber.Router) {
	app.Post("/", CastRedisVote)
	app.Get("/history/:electionID", GetRedisVotesByElectionIDHandler)
}

var Hub *ws.Hub

func CastRedisVote(c *fiber.Ctx) error {
	var votes []operations.RedisVote
	if err := c.BodyParser(&votes); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID")
	}

	if len(votes) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "No votes submitted")
	}

	// Check if already voted
	exists, err := operations.HasAlreadyVotedRedis(voterCardID, votes[0].ElectionID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check vote status")
	}
	if exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "You have already voted in this election",
		})
	}

	for _, vote := range votes {
		vote.VoterCardID = voterCardID
		if err := operations.CacheVote(&vote); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to cache vote")
		}
		msg, _ := json.Marshal(vote)
		Hub.Broadcast(msg)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully cast all votes",
	})
}

func GetRedisVotesByElectionIDHandler(c *fiber.Ctx) error {
	electionID, err := strconv.Atoi(c.Params("electionID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid election ID")
	}

	votes, err := operations.GetRedisVotesByElectionID(electionID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch votes")
	}

	return c.JSON(votes)
}
