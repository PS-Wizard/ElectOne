package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type VoterCard struct {
	VoterCardID string `json:"voter_card_id"`
	CitizenID   string `json:"citizenship_id"`
	Location    string `json:"location"`
	Privileges  string `json:"voting_privileges"`
}

func CreateVoterCard(voterCard *VoterCard) (string, error) {
	query := `
        INSERT INTO VoterCard (VoterCardID, CitizenID, Location, ElectionType)
        VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, voterCard.VoterCardID, voterCard.CitizenID, voterCard.Location, voterCard.Privileges)
	if err != nil {
		return "", err
	}
	return voterCard.VoterCardID, nil
}

func UpdateVoterCard(voterCardID string, updatedVoterCard *VoterCard) error {
	query := `
        UPDATE VoterCard
        SET CitizenID = ?, Location = ?, ElectionType = ?
        WHERE VoterCardID = ?`
	_, err := db.DB.Exec(query, updatedVoterCard.CitizenID, updatedVoterCard.Location, updatedVoterCard.Privileges, voterCardID)
	return err
}

func GetVoterCards(limit, offset int) ([]VoterCard, error) {
	query := `
        SELECT VoterCardID, CitizenID, Location, ElectionType
        FROM VoterCard
        ORDER BY VoterCardID DESC
        LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var voterCards []VoterCard
	for rows.Next() {
		var voterCard VoterCard
		if err := rows.Scan(&voterCard.VoterCardID, &voterCard.CitizenID, &voterCard.Location, &voterCard.Privileges); err != nil {
			return nil, err
		}
		voterCards = append(voterCards, voterCard)
	}
	return voterCards, nil
}

func GetVoterCardByID(voterCardID string) (*VoterCard, error) {
	query := `
        SELECT VoterCardID, CitizenID, Location, ElectionType
        FROM VoterCard
        WHERE VoterCardID = ?`
	var voterCard VoterCard
	err := db.DB.QueryRow(query, voterCardID).Scan(&voterCard.VoterCardID, &voterCard.CitizenID, &voterCard.Location, &voterCard.Privileges)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No VoterCard With That ID")
		}
		return nil, err
	}
	return &voterCard, nil
}

func DeleteVoterCard(voterCardID string) error {
	query := `DELETE FROM VoterCard WHERE VoterCardID = ?`
	_, err := db.DB.Exec(query, voterCardID)
	return err
}
