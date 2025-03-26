package candidates

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

func getCandidateDetail(candidateID string) (*Candidate, error) {
	var candidate Candidate
	query := "SELECT candidateID, citizenID, post, electionID, GroupName FROM candidates WHERE candidateID = ?"
	row := api.DB.QueryRow(query, candidateID)
	err := row.Scan(&candidate.CandidateID, &candidate.CitizenID, &candidate.Post, &candidate.ElectionID, &candidate.Group)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("candidate not found")
		}
		return nil, fmt.Errorf("failed to fetch candidate: %v", err)
	}
	return &candidate, nil
}

func getCandidatesPaginated(offset int) ([]Candidate, error) {
	var candidates []Candidate
	query := "SELECT candidateID, citizenID, post, electionID, GroupName FROM candidates LIMIT 10 OFFSET ?"
	rows, err := api.DB.Query(query, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch candidates: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var candidate Candidate
		err := rows.Scan(&candidate.CandidateID, &candidate.CitizenID, &candidate.Post, &candidate.ElectionID, &candidate.Group)
		if err != nil {
			return nil, fmt.Errorf("failed to scan candidate: %v", err)
		}
		candidates = append(candidates, candidate)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return candidates, nil
}

func createCandidate(candidate Candidate) error {
	query := "INSERT INTO candidates (citizenID, post, electionID, GroupName) VALUES (?, ?, ?, ?)"
	_, err := api.DB.Exec(query, candidate.CitizenID, candidate.Post, candidate.ElectionID, candidate.Group)
	if err != nil {
		return fmt.Errorf("failed to insert candidate: %v", err)
	}
	return nil
}

func updateCandidate(candidateID string, candidate Candidate) error {
	query := "UPDATE candidates SET citizenID = ?, post = ?, electionID = ?, GroupName = ? WHERE candidateID = ?"
	_, err := api.DB.Exec(query, candidate.CitizenID, candidate.Post, candidate.ElectionID, candidate.Group, candidateID)
	if err != nil {
		return fmt.Errorf("failed to update candidate: %v", err)
	}
	return nil
}

func deleteCandidate(candidateID string) error {
	query := "DELETE FROM candidates WHERE candidateID = ?"
	_, err := api.DB.Exec(query, candidateID)
	if err != nil {
		return fmt.Errorf("failed to delete candidate: %v", err)
	}
	return nil
}
