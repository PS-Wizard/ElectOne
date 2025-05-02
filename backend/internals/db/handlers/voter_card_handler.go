package handlers

import (
	"strconv"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterVoterCardRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateVoterCardHandler)
	router.Get("/:id", GetVoterCardByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateVoterCardHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteVoterCardHandler)
	router.Get("/", middlewares.RequireAdmin, ListVoterCardsHandler)
}

func CreateVoterCardHandler(c *fiber.Ctx) error {
	var voterCard operations.VoterCard
	if err := c.BodyParser(&voterCard); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid voter card input")
	}

	id, err := operations.CreateVoterCard(&voterCard)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(fiber.Map{"voter_card_id": id})
}

func GetVoterCardByIDHandler(c *fiber.Ctx) error {
	var id string

	token := c.Locals("token").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
	}

	role := claims["role"]

	if role == "user" {
		id = claims["voter_card"].(string)
	} else {
		id = c.Params("id")
	}
	voterCard, err := operations.GetVoterCardByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(voterCard)
}

func UpdateVoterCardHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var updated operations.VoterCard

	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateVoterCard(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update voter card")
	}

	return c.SendStatus(fiber.StatusOK)
}

func DeleteVoterCardHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := operations.DeleteVoterCard(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete voter card")
	}

	return c.SendStatus(fiber.StatusOK)
}

func ListVoterCardsHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	voterCards, err := operations.GetVoterCards(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch voter cards")
	}

	return c.JSON(voterCards)
}
