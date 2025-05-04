package statistics

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func GetElectionHistory(c *fiber.Ctx) error {
	electionID := c.Query("election_id")
	if electionID == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Election ID is required")
	}

	posts := []string{"Mayor", "President", "Senator"} 

	history := make(map[string][]map[string]any) // To store history for each post

	for _, post := range posts {
		historyKey := fmt.Sprintf("history:election:%s:%s", electionID, post)

		// Fetch the sorted set from Redis (voting history)
		results, err := db.RDB.ZRangeByScoreWithScores(context.Background(), historyKey, &redis.ZRangeBy{
			Min: "-inf", Max: "+inf",
		}).Result()
		if err != nil {
			log.Println("Failed to fetch voting history:", err)
			continue
		}

		// Process results into a list of entries
		postHistory := []map[string]any{}
		for _, result := range results {
			entry := result.Member.(string) // Format: candidateID-timestamp
			parts := strings.Split(entry, "-")
			if len(parts) == 2 {
				candidateID := parts[0]
				timestamp := parts[1]

				// Add entry to the history for the post
				postHistory = append(postHistory, map[string]any{
					"candidate_id": candidateID,
					"timestamp":    timestamp,
				})
			}
		}

		// Store the history for this post
		history[post] = postHistory
	}

	// Send the collected history as JSON
	return c.JSON(history)
}
