package redis

import (
	"fmt"
	"log"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

func incrementVote(candidateID, electionID string) error {
	key := fmt.Sprintf("votes:%s:%s", candidateID, electionID)
	log.Printf("Incrementing vote for key: %s", key) // Log the key being used
	cmd := api.RDB.Incr(api.CTX, key)
	num, err := cmd.Result()
	log.Printf("Num: %d", num) // Log the key being used
	if err != nil {
		log.Printf("Failed To Increment Vote: %v", err)
		return err
	}
	msg := fmt.Sprintf("Vote Count Updated For Election %s, Candidate %s: %d", electionID, candidateID, num)
	err = api.RDB.Publish(api.CTX, "voteUpdates", msg).Err()
	if err != nil {
		log.Printf("Failed to publish vote count update: %v", err)
		return err
	}
	fmt.Printf("Publisehd vote update:")
	return nil
}
