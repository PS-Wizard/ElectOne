package candidates

import (
	"fmt"
	"strings"

	"github.com/PS-Wizard/VotingSystem/db"
)

type Candidate struct {
	CitizenID  uint   `json:"citizen_id"`
	ElectionID uint   `json:"election_id"`
	Post       string `json:"post"`
	TotalVotes uint   `json:"total_votes"`
}

func (r *Candidate) CreateRunner() error {
	query := `INSERT INTO runners (citizen_id, election_id, post, total_votes) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, r.CitizenID, r.ElectionID, r.Post, r.TotalVotes)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("runner for this election already exists")
		}
		return fmt.Errorf("failed to insert runner into database: %v", err)
	}
	fmt.Println("Candidate Successfully Created!")
	return nil
}

func GetRunners() ([]Candidate, error) {
	rows, err := db.DB.Query("SELECT citizen_id, election_id, post, total_votes FROM runners")
	if err != nil {
		return nil, fmt.Errorf("failed to query runners: %v", err)
	}
	defer rows.Close()

	var runners []Candidate
	for rows.Next() {
		var runner Candidate
		err := rows.Scan(&runner.CitizenID, &runner.ElectionID, &runner.Post, &runner.TotalVotes)
		if err != nil {
			return nil, fmt.Errorf("failed to scan runner data: %v", err)
		}
		runners = append(runners, runner)
	}

	return runners, nil
}
