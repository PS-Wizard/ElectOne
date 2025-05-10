package operations

import (
	"github.com/PS-Wizard/Electone/internals/db"
)

type Vote struct {
	VoteID      int    `json:"vote_id,omitempty"`
	VoterCardID string `json:"voter_card_id,omitempty"`
	ElectionID  int    `json:"election_id"`
	CandidateID int    `json:"candidate_id"`
	CreatedAt   string `json:"created_at"`
}

func CreateVote(vote *Vote) (int64, error) {
	query := `INSERT INTO Votes (VoterCardID, ElectionID, CandidateID) VALUES (?, ?, ?)`
	result, err := db.DB.Exec(query, vote.VoterCardID, vote.ElectionID, vote.CandidateID)
	if err != nil {
		return 0, err
	}
	// BroadcastVoteUpdate(*vote)
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

func GetVotesByElectionID(electionID int, limit, offset int) ([]Vote, error) {
	query := `SELECT VoteID, VoterCardID, ElectionID, CandidateID, created_at FROM Votes WHERE ElectionID = ? ORDER BY VoteID DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, electionID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var votes []Vote
	for rows.Next() {
		var vote Vote
		if err := rows.Scan(&vote.VoteID, &vote.VoterCardID, &vote.ElectionID, &vote.CandidateID, &vote.CreatedAt); err != nil {
			return nil, err
		}
		votes = append(votes, vote)
	}
	return votes, nil
}

func HasAlreadyVoted(voterCardID string, electionID int) (bool, error) {
	query := `SELECT COUNT(1) FROM Votes WHERE VoterCardID = ? AND ElectionID = ?`
	var count int
	err := db.DB.QueryRow(query, voterCardID, electionID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
