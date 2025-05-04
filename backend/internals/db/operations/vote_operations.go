package operations

import (
	"github.com/PS-Wizard/Electone/internals/db"
)

type Vote struct {
	VoteID      int    `json:"vote_id,omitempty"`
	VoterCardID string `json:"voter_card_id,omitempty"`
	ElectionID  int    `json:"election_id"`
	CandidateID int    `json:"candidate_id"`
}

func CreateVote(vote *Vote) (int64, error) {
	query := `INSERT INTO Votes (VoterCardID, ElectionID, CandidateID) VALUES (?, ?, ?)`
	result, err := db.DB.Exec(query, vote.VoterCardID, vote.ElectionID, vote.CandidateID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func GetVotes(limit, offset int) ([]Vote, error) {
	query := `SELECT VoteID, VoterCardID, ElectionID, CandidateID FROM Votes ORDER BY VoteID DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []Vote
	for rows.Next() {
		var vote Vote
		if err := rows.Scan(&vote.VoteID, &vote.VoterCardID, &vote.ElectionID, &vote.CandidateID); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func GetVotesByVoterCard(voterCardID string) ([]Vote, error) {
	query := `SELECT VoteID, VoterCardID, ElectionID, CandidateID FROM Votes WHERE VoterCardID = ?`
	rows, err := db.DB.Query(query, voterCardID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []Vote
	for rows.Next() {
		var vote Vote
		if err := rows.Scan(&vote.VoteID, &vote.VoterCardID, &vote.ElectionID, &vote.CandidateID); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func DeleteVote(voteID int) error {
	query := `DELETE FROM Votes WHERE VoteID = ?`
	_, err := db.DB.Exec(query, voteID)
	return err
}
