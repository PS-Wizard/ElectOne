package handlers

import (
	"fmt"
	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func RegisterElectionRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireAdmin, CreateElectionHandler)
	router.Get("/:id", GetElectionByIDHandler)
	router.Put("/:id", middlewares.RequireAdmin, UpdateElectionHandler)
	router.Delete("/:id", middlewares.RequireAdmin, DeleteElectionHandler)
	router.Get("/", ListElectionsHandler)
	router.Get("/candidates/:electionID", GetCandidatesByElectionIDHandler)
}

func CreateElectionHandler(c *fiber.Ctx) error {
	var election operations.Election
	if err := c.BodyParser(&election); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	fmt.Println("Got Request: ", election)
	id, err := operations.CreateElection(&election)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create election")
	}

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
		// Get voter card ID from claims
		voterCardID, ok := claims["voter_card"].(string)
		if !ok {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID in token")
		}

		// Fetch the voter card data
		voterCard, err := operations.GetVoterCardByID(voterCardID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch voter card")
		}

		// Check if the election's location matches the voter's location
		if election.Location != voterCard.Location {
			return fiber.NewError(fiber.StatusForbidden, "You are not allowed to access this election due to location mismatch")
		}

		// Check if the election type matches the voter's privileges
		fmt.Println("Location: ", election.Location, voterCard.Location)
		fmt.Println("Privileges: ", election.Type, voterCard.Privileges)
		if election.Type != voterCard.Privileges {
			return fiber.NewError(fiber.StatusForbidden, "You are not allowed to vote in this election due to privilege mismatch")
		}
	}

	// Return the election details as JSON if all checks pass
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

	return c.SendStatus(fiber.StatusOK)
}

func ListElectionsHandler(c *fiber.Ctx) error {
	token := c.Locals("token").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	role := claims["role"]

	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	offset, _ := strconv.Atoi(c.Query("offset", "0"))

	// If the role is admin, return all elections without any restrictions
	if role == "admin" {
		elections, err := operations.GetElections(limit, offset)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(elections)
	}

	// For users, apply location-based and privilege-based validation
	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID in token")
	}

	// Fetch the voter card data
	voterCard, err := operations.GetVoterCardByID(voterCardID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch voter card")
	}

	// Fetch elections based on the voter's location
	elections, err := operations.GetElectionsByLocation(voterCard.Location, limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Filter elections by checking if the user's privileges match the election type
	var filteredElections []operations.Election
	for _, election := range elections {
		if election.Type == voterCard.Privileges {
			filteredElections = append(filteredElections, election)
		}
	}

	// Return the filtered elections
	return c.JSON(filteredElections)
}

func GetCandidatesByElectionIDHandler(c *fiber.Ctx) error {
	electionID, err := strconv.Atoi(c.Params("electionID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("Invalid ElectionID"))
	}
	candidates, err := operations.GetCandidatesByElectionID(electionID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch candidates")
	}

	return c.JSON(candidates)
}
