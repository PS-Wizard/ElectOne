package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/gofiber/fiber/v2"
)

func SyncRedisToTursoHandler(c *fiber.Ctx) error {
	electionIDStr := c.Params("electionID")
	electionID, err := strconv.Atoi(electionIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid election ID format"})
	}

	fmt.Println("Starting Redis to TursoDB sync process ...")

	// 1. Fetch votes from Redis
	redisVotes, err := operations.GetRedisVotesByElectionID(electionID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch votes from Redis"})
	}

	if len(redisVotes) == 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":            "No votes in Redis to sync for this election.",
			"total_redis_votes":  0,
			"synced_to_turso":    0,
			"skipped_duplicates": 0,
			"errors_during_sync": 0,
		})
	}

	fmt.Println("Fetched votes for election: ", electionID)

	var syncedCount, skippedDuplicatesCount, errorCount int

	for _, redisVote := range redisVotes {
		t := time.Unix(redisVote.CastUnixTime, 0)

		tursoVote := operations.Vote{
			ElectionID:  redisVote.ElectionID,
			VoterCardID: redisVote.VoterCardID,
			CandidateID: redisVote.CandidateID,
			CreatedAt:   t.Format("2006-01-02 15:04:05"), 
		}

		_, err = operations.CreateVote(&tursoVote)
		if err != nil {
			fmt.Println(err)
			errorCount++
			continue 
		}
		syncedCount++
	}

	summary := fiber.Map{
		"message":            "Sync process completed for election.",
		"election_id":        electionID,
		"total_redis_votes":  len(redisVotes),
		"synced_to_turso":    syncedCount,
		"skipped_duplicates": skippedDuplicatesCount,
		"errors_during_sync": errorCount,
		"redis_list_cleared": errorCount == 0 && len(redisVotes) > 0,
	}
	return c.Status(fiber.StatusOK).JSON(summary)
}

func RegisterSyncRoutes(router fiber.Router) {
	router.Post("/election/:electionID/redis-to-turso", SyncRedisToTursoHandler)
}
