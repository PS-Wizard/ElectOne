package elections

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

func getElectionDetail(id string) (*Election, error) {
	var election Election
	query := "SELECT electionID, title, startDate, endDate, votingRestrictions FROM elections WHERE electionID = ?"
	row := api.DB.QueryRow(query, id)
	err := row.Scan(&election.ElectionID, &election.Title, &election.StartDate, &election.EndDate, &election.VotingRestrictions)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("election not found")
		}
		return nil, fmt.Errorf("failed to fetch election: %v", err)
	}
	return &election, nil
}

func getElectionsPaginated(offset int) ([]Election, error) {
	var elections []Election
	query := "SELECT electionID, title, startDate, endDate, votingRestrictions FROM elections LIMIT 10 OFFSET ?"
	rows, err := api.DB.Query(query, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch elections: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var election Election
		err := rows.Scan(&election.ElectionID, &election.Title, &election.StartDate, &election.EndDate, &election.VotingRestrictions)
		if err != nil {
			return nil, fmt.Errorf("failed to scan election: %v", err)
		}
		elections = append(elections, election)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}
	return elections, nil
}

func createElection(details Election) error {
	query := "INSERT INTO elections (title, startDate, endDate, votingRestrictions) VALUES (?, ?, ?, ?)"
	_, err := api.DB.Exec(query, details.Title, details.StartDate, details.EndDate, details.VotingRestrictions)
	if err != nil {
		return fmt.Errorf("failed to insert election details: %v", err)
	}
	return nil
}

func updateElectionDetail(id string, newDetail Election) error {
	query := "UPDATE elections SET title = ?, startDate = ?, endDate = ?, votingRestrictions = ? WHERE electionID = ?"
	_, err := api.DB.Exec(query, newDetail.Title, newDetail.StartDate, newDetail.EndDate, newDetail.VotingRestrictions, id)
	if err != nil {
		return fmt.Errorf("failed to update election details: %v", err)
	}
	return nil
}

func deleteElection(id string) error {
	query := "DELETE FROM elections WHERE electionID = ?"
	_, err := api.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete election: %v", err)
	}
	return nil
}
