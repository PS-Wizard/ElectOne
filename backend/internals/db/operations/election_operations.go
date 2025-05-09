package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type Election struct {
	ElectionID  int    `json:"election_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Location    string `json:"location"`
}

type CandidateInfo struct {
	CandidateID int    `json:"candidate_id"`
	ProfilePath string `json:"candidate_photo"`
	Bio         string `json:"candidate_bio"`
	Post        string `json:"candidate_post"`
	Party       string `json:"candidate_party"`
	Name        string `json:"candidate_name"`
}

func CreateElection(election *Election) (int, error) {
	query := `
		INSERT INTO Election (Name, Description, StartDate, EndDate, Location)
		VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, election.Name, election.Description, election.StartDate, election.EndDate, election.Location)
	if err != nil {
		return 0, err
	}

	var electionID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&electionID)
	if err != nil {
		return 0, err
	}

	return electionID, nil
}

func UpdateElection(electionID int, updatedElection *Election) error {
	query := `
		UPDATE Election
		SET Name = ?, Description = ?, StartDate = ?, EndDate = ?, Location = ?
		WHERE ElectionID = ?`
	_, err := db.DB.Exec(query, updatedElection.Name, updatedElection.Description, updatedElection.StartDate, updatedElection.EndDate, updatedElection.Location, electionID)
	return err
}

func GetElections(limit, offset int) ([]Election, error) {
	query := `
		SELECT ElectionID, Name, Description, StartDate, EndDate, Location
		FROM Election
		ORDER BY ElectionID DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var elections []Election
	for rows.Next() {
		var election Election
		if err := rows.Scan(&election.ElectionID, &election.Name, &election.Description, &election.StartDate, &election.EndDate, &election.Location); err != nil {
			return nil, err
		}
		elections = append(elections, election)
	}
	return elections, nil
}

func GetElectionByID(electionID int) (*Election, error) {
	query := `
		SELECT ElectionID, Name, Description, StartDate, EndDate, Location
		FROM Election
		WHERE ElectionID = ?`
	var election Election
	err := db.DB.QueryRow(query, electionID).Scan(&election.ElectionID, &election.Name, &election.Description, &election.StartDate, &election.EndDate, &election.Location)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Election With That ID")
		}
		return nil, err
	}
	return &election, nil
}

func DeleteElection(electionID int) error {
	query := `DELETE FROM Election WHERE ElectionID = ?`
	_, err := db.DB.Exec(query, electionID)
	return err
}

func GetElectionsByLocation(location string, limit, offset int) ([]Election, error) {
	query := `
		SELECT ElectionID, Name, Description, StartDate, EndDate, Location
		FROM Election
		WHERE Location = ?
		ORDER BY ElectionID DESC
		LIMIT ? OFFSET ?
	`

	rows, err := db.DB.Query(query, location, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var elections []Election
	for rows.Next() {
		var e Election
		if err := rows.Scan(&e.ElectionID, &e.Name, &e.Description, &e.StartDate, &e.EndDate, &e.Location); err != nil {
			return nil, err
		}
		elections = append(elections, e)
	}

	return elections, nil
}

func GetCandidatesByElectionID(electionID int) ([]CandidateInfo, error) {
	query := `
		SELECT CandidateID, Name, ProfilePath, Bio, Post, Party
		FROM Candidate
		WHERE ElectionID = ?
	`
	rows, err := db.DB.Query(query, electionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var candidates []CandidateInfo
	for rows.Next() {
		var c CandidateInfo
		if err := rows.Scan(&c.CandidateID, &c.Name, &c.ProfilePath, &c.Bio, &c.Post, &c.Party); err != nil {
			return nil, err
		}
		candidates = append(candidates, c)
	}
	return candidates, nil

}
