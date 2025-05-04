package handlers

import (
	"fmt"
	"strconv"
	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RegisterElectionRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateElectionHandler)
	router.Get("/:id", GetElectionByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateElectionHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteElectionHandler)
	router.Get("/", ListElectionsHandler)
}

func CreateElectionHandler(c *fiber.Ctx) error {
	var election operations.Election
	if err := c.BodyParser(&election); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	id, err := operations.CreateElection(&election)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create election")
	}

	fmt.Println("New Election Created, removing cache")
	return c.JSON(fiber.Map{"election_id": id})
}

func GetElectionByIDHandler(c *fiber.Ctx) error {
	token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	role := claims["role"]
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	election, err := operations.GetElectionByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if role == "user" {
		voterCardID, ok := claims["voter_card"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID in token")
		}

		voterCard, err := operations.GetVoterCardByID(voterCardID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch voter card")
		}

		if election.Location != voterCard.Location {
			return fiber.NewError(fiber.StatusForbidden, "You are not allowed to access this election")
		}
	}

	return c.JSON(election)
}

func UpdateElectionHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	var updated operations.Election
	if err := c.BodyParser(&updated); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := operations.UpdateElection(id, &updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fmt.Println("Eleciton Updated, Refreshing Cache")
	return c.SendStatus(fiber.StatusOK)
}

func DeleteElectionHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid ID")
	}

	if err := operations.DeleteElection(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete election")
	}

	fmt.Println("Eleciton Updated, Refreshing Cache")
	return c.SendStatus(fiber.StatusOK)
}

func ListElectionsHandler(c *fiber.Ctx) error {
	fmt.Println("Ive been called")
	token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"]

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	if role == "admin" {
		elections, err := operations.GetElections(limit, offset)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(elections)
	}

	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID in token")
	}

	voterCard, err := operations.GetVoterCardByID(voterCardID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch voter card")
	}

	elections, err := operations.GetElectionsByLocation(voterCard.Location, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(elections)
}
