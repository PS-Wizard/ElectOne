package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type Vote struct {
	VoteID      int    `json:"vote_id"`
	VoterCardID string `json:"voter_card_id"`
	ElectionID  int    `json:"election_id"`
	VotedFor    string `json:"voted_for"` // Json
	VoteDate    string `json:"vote_date"`
}

func CreateVote(vote *Vote) (int, error) {
	query := `
		INSERT INTO Vote (VoterCardID, ElectionID, VotedFor, VoteDate)
		VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, vote.VoterCardID, vote.ElectionID, vote.VotedFor, vote.VoteDate)
	if err != nil {
		return 0, err
	}

	var voteID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&voteID)
	if err != nil {
		return 0, err
	}

	return voteID, nil
}

func GetVotes(limit, offset int) ([]Vote, error) {
	query := `
		SELECT VoteID, VoterCardID, ElectionID, VotedFor, VoteDate
		FROM Vote
		ORDER BY VoteID DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []Vote
	for rows.Next() {
		var vote Vote
		if err := rows.Scan(&vote.VoteID, &vote.VoterCardID, &vote.ElectionID, &vote.VotedFor, &vote.VoteDate); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func GetVoteByID(voteID int) (*Vote, error) {
	query := `
		SELECT VoteID, VoterCardID, ElectionID, VotedFor, VoteDate
		FROM Vote
		WHERE VoteID = ?`
	var vote Vote
	err := db.DB.QueryRow(query, voteID).Scan(&vote.VoteID, &vote.VoterCardID, &vote.ElectionID, &vote.VotedFor, &vote.VoteDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Vote With That ID")
		}
		return nil, err
	}
	return &vote, nil
}
