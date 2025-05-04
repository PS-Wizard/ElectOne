package statistics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

type VoteReq struct {
	ElectionID  int            `json:"election_id"`
	Candidates  map[string]any `json:"candidates"`
	VoterCardID string         `json:"voter_card_id,omitempty"`
}

func RegisterVoteRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireUser, HandleVote)

}

func HandleVote(c *fiber.Ctx) error {
	var vote VoteReq
	// Extract token and claims
	token := c.Locals("token").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid claims")
	}

	// Extract VoterCardID from token claims
	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("No voter card found in token")
	}
	vote.VoterCardID = voterCardID

	// Parse JSON body into VoteReq struct
	if err := c.BodyParser(&vote); err != nil {
		return c.Status(400).SendString("Invalid vote payload")
	}

	// Prevent double voting
	// voteLockKey := fmt.Sprintf("voted:election:%d:voter:%s", vote.ElectionID, vote.VoterCardID)
	// set, err := db.RDB.SetNX(context.Background(), voteLockKey, "1", 0).Result()
	// if err != nil {
	// 	return c.Status(500).SendString("Redis error")
	// }
	// if !set {
	// 	return c.Status(fiber.StatusForbidden).SendString("You have already voted in this election")
	// }

	// Tally votes
	for post, candidateID := range vote.Candidates {
		// Construct voteKey for each post and candidate
		voteKey := fmt.Sprintf("election:%d:%s:%s:votes", vote.ElectionID, post, candidateID)

		// Increment vote count using Redis
		if err := db.RDB.Incr(context.Background(), voteKey).Err(); err != nil {
			return c.Status(500).SendString("Failed to record vote")
		}

		timestamp := float64(time.Now().Unix()) // or .UnixMilli() if you want ms precision

		historyKey := fmt.Sprintf("history:election:%d:%s", vote.ElectionID, post)

		// Unique ID for each vote entry (timestamp + random suffix to ensure uniqueness)
		entry := fmt.Sprintf("%v-%v", candidateID, timestamp)

		// Add to Redis sorted set for history with a unique member
		err := db.RDB.ZAdd(context.Background(), historyKey, redis.Z{
			Score:  timestamp,
			Member: entry,
		}).Err()
		if err != nil {
			log.Println("failed to record vote history:", err)
		}
	}

	// Publish update to Redis channel
	update := map[string]any{
		"election_id":   vote.ElectionID,
		"candidates":    vote.Candidates,
		"voter_card_id": vote.VoterCardID,
	}
	msg, _ := json.Marshal(update)
	if err := db.RDB.Publish(context.Background(), "vote_updates", msg).Err(); err != nil {
		log.Println("failed to publish vote update:", err)
	}

	return c.SendString("Vote cast successfully")
}
