package realtime

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type VoteReq struct {
	ElectionID   int      `json:"election_id"`
	CandidateIDs []string `json:"candidate_ids"`
	VoterCardID  string   `json:"voter_card_id,omitempty"`
}

func RegisterVoteRoutes(router fiber.Router) {
	router.Post("/", middlewares.RequireUser, HandleVote)

}

func HandleVote(c *fiber.Ctx) error {
	var vote VoteReq

	token := c.Locals("token").(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid claims")
	}

	voterCardID, ok := claims["voter_card"].(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("No voter card found in token")
	}
	vote.VoterCardID = voterCardID

	// Parse JSON body
	if err := c.BodyParser(&vote); err != nil {
		return c.Status(400).SendString("Invalid vote payload")
	}

	// Prevent double voting
	voteLockKey := fmt.Sprintf("voted:election:%d:voter:%s", vote.ElectionID, vote.VoterCardID)
	set, err := db.RDB.SetNX(context.Background(), voteLockKey, "1", 0).Result()
	if err != nil {
		return c.Status(500).SendString("Redis error")
	}
	if !set {
		return c.Status(fiber.StatusForbidden).SendString("You have already voted in this election")
	}

	// Tally votes
	for _, candidateID := range vote.CandidateIDs {
		voteKey := fmt.Sprintf("election:%d:candidate:%s:votes", vote.ElectionID, candidateID)
		if err := db.RDB.Incr(context.Background(), voteKey).Err(); err != nil {
			return c.Status(500).SendString("Failed to record vote")
		}
	}

	update := map[string]any{
		"election_id":     vote.ElectionID,
		"candidate_ids":   vote.CandidateIDs,
		"voter_card_id":   vote.VoterCardID,
		"vote_count_keys": fmt.Sprintf("election:%d:*:votes", vote.ElectionID),
	}
	msg, _ := json.Marshal(update)
	if err := db.RDB.Publish(context.Background(), "vote_updates", msg).Err(); err != nil {
		log.Println("failed to publish vote update:", err)
	}

	return c.SendString("Vote cast successfully")
}
