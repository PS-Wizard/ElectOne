package operations

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/PS-Wizard/Electone/internals/db"
)

type RedisVote struct {
	ElectionID   int    `json:"election_id"`
	VoterCardID  string `json:"voter_card_id"`
	CandidateID  int    `json:"candidate_id"`
	CastUnixTime int64  `json:"cast_unix_time"`
}

func getRedisKey(electionID int) string {
	return fmt.Sprintf("votes:election:%d", electionID)
}

func CacheVote(vote *RedisVote) error {
	key := getRedisKey(vote.ElectionID)
	vote.CastUnixTime = time.Now().Unix()

	data, err := json.Marshal(vote)
	if err != nil {
		return err
	}

	return db.RDB.RPush(context.Background(), key, data).Err()
}

func HasAlreadyVotedRedis(voterCardID string, electionID int) (bool, error) {
	key := getRedisKey(electionID)
	votes, err := db.RDB.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return false, err
	}

	for _, v := range votes {
		var vote RedisVote
		if err := json.Unmarshal([]byte(v), &vote); err == nil {
			if vote.VoterCardID == voterCardID {
				return true, nil
			}
		}
	}

	return false, nil
}
func GetRedisVotesByElectionID(electionID int) ([]RedisVote, error) {
	key := getRedisKey(electionID)

	votesData, err := db.RDB.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	votes := make([]RedisVote, 0, len(votesData))
	for _, v := range votesData {
		var vote RedisVote
		if err := json.Unmarshal([]byte(v), &vote); err != nil {
			continue
		}
		votes = append(votes, vote)
	}

	return votes, nil
}
