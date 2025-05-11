package handlers

import (
	"fmt"
	"strconv"
	"strings"

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

		// Check location match
		locationMatch := false
		allowedLocations := strings.Split(election.Location, ",")
		for _, loc := range allowedLocations {
			if strings.TrimSpace(loc) == voterCard.Location {
				locationMatch = true
				break
			}
		}
		if !locationMatch {
			return fiber.NewError(fiber.StatusForbidden, "You are not allowed to access this election due to location mismatch")
		}

		// Check privilege match
		privilegeMatch := false
		voterPrivileges := strings.Split(voterCard.Privileges, ",")
		for _, priv := range voterPrivileges {
			if strings.TrimSpace(priv) == strings.TrimSpace(election.Type) {
				privilegeMatch = true
				break
			}
		}
		if !privilegeMatch {
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

	// Admin: return all elections
	if role == "admin" {
		elections, err := operations.GetElections(limit, offset)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(elections)
	}

	// User: apply voterCard filters
	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid voter card ID in token")
	}

	voterCard, err := operations.GetVoterCardByID(voterCardID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not fetch voter card")
	}

	// Fetch ALL elections
	elections, err := operations.GetElections(limit, offset)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// Filter by location and privilege
	var filteredElections []operations.Election
	// voterLocations := []string{strings.TrimSpace(voterCard.Location)}
	voterPrivileges := strings.Split(voterCard.Privileges, ",")

	for _, election := range elections {
		allowedLocations := strings.Split(election.Location, ",")
		requiredType := strings.TrimSpace(election.Type)

		locationMatch := false
		for _, loc := range allowedLocations {
			if strings.TrimSpace(loc) == voterCard.Location {
				locationMatch = true
				break
			}
		}

		privilegeMatch := false
		for _, priv := range voterPrivileges {
			if strings.TrimSpace(priv) == requiredType {
				privilegeMatch = true
				break
			}
		}

		if locationMatch && privilegeMatch {
			filteredElections = append(filteredElections, election)
		}
	}

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
