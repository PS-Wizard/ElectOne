package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
)

func RegisterVoteRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateVoteHandler)
	router.Get("/:id", middlewares.RequireAdmin, GetVoteByIDHandler)
	router.Get("/", middlewares.RequireAdmin, ListVotesHandler)
}

func CreateVoteHandler(c *fiber.Ctx) error {
	var vote operations.Vote
	if err := c.BodyParser(&vote); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid vote input")
	}

	id, err := operations.CreateVote(&vote)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"vote_id": id})
}

func GetVoteByIDHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid vote ID")
	}

	vote, err := operations.GetVoteByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(vote)
}

func ListVotesHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	votes, err := operations.GetVotes(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch votes")
	}

	return c.JSON(votes)
}
