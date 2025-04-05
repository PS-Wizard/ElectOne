package redis

import (
	"fmt"
	"github.com/PS-Wizard/ElectOneAPI/api"
)

func incrementVote(candidateID, electionID string) error {
	key := fmt.Sprintf("votes:%s:%s", candidateID, electionID)
	cmd := api.RDB.Incr(api.CTX, key)
	num, err := cmd.Result()
	if err != nil {
		return err
	}

	msg := fmt.Sprintf(`{"electionID":"%s", "candidateID":"%s", "votes":%d}`, electionID, candidateID, num)
	api.RDB.Publish(api.CTX, "voteUpdates", msg)

	return nil
}
