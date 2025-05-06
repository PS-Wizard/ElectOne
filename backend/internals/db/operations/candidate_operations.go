package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type Candidate struct {
	CandidateID int    `json:"candidate_id"`
	CitizenID   string `json:"citizen_id"`
	ElectionID  int    `json:"election_id"`
	ProfilePath string `json:"candidate_photo"`
	Bio         string `json:"candidate_bio"`
	Post        string `json:"candidate_post"`
	Party       string `json:"candidate_party"`
	Name        string `json:"candidate_name"`
}

func CreateCandidate(candidate *Candidate) (int, error) {
	query := `
		INSERT INTO Candidate (CitizenID, ElectionID, ProfilePath, Bio, Post, Party, Name)
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, candidate.CitizenID, candidate.ElectionID, candidate.ProfilePath, candidate.Bio, candidate.Post, candidate.Party, candidate.Name)
	if err != nil {
		return 0, err
	}

	var candidateID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&candidateID)
	if err != nil {
		return 0, err
	}

	return candidateID, nil
}

func UpdateCandidate(candidateID int, updatedCandidate *Candidate) error {
	query := `
		UPDATE Candidate
		SET CitizenID = ?, ElectionID = ?, ProfilePath = ?, Bio = ?, Post = ?, Party = ?, Name = ?
		WHERE CandidateID = ?`
	_, err := db.DB.Exec(query, updatedCandidate.CitizenID, updatedCandidate.ElectionID, updatedCandidate.ProfilePath, updatedCandidate.Bio, updatedCandidate.Post, updatedCandidate.Party, updatedCandidate.Name, candidateID)
	return err
}

func GetCandidates(limit, offset int) ([]Candidate, error) {
	query := `
		SELECT CandidateID, CitizenID, ElectionID, ProfilePath, Bio, Post, Party, Name
		FROM Candidate
		ORDER BY CandidateID DESC
		LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var candidates []Candidate
	for rows.Next() {
		var candidate Candidate
		if err := rows.Scan(&candidate.CandidateID, &candidate.CitizenID, &candidate.ElectionID, &candidate.ProfilePath, &candidate.Bio, &candidate.Post, &candidate.Party, &candidate.Name); err != nil {
			return nil, err
		}
		candidates = append(candidates, candidate)
	}
	return candidates, nil
}

func GetCandidateByID(candidateID int) (*Candidate, error) {
	query := `
		SELECT CandidateID, CitizenID, ElectionID, ProfilePath, Bio, Post, Party, Name
		FROM Candidate
		WHERE CandidateID = ?`
	var candidate Candidate
	err := db.DB.QueryRow(query, candidateID).Scan(&candidate.CandidateID, &candidate.CitizenID, &candidate.ElectionID, &candidate.ProfilePath, &candidate.Bio, &candidate.Post, &candidate.Party, &candidate.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Candidate With That ID")
		}
		return nil, err
	}
	return &candidate, nil
}

func DeleteCandidate(candidateID int) error {
	query := `DELETE FROM Candidate WHERE CandidateID = ?`
	_, err := db.DB.Exec(query, candidateID)
	return err
}
