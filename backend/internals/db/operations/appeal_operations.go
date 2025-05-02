package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

// represents the temporary storage of user registration data.
type Appeal struct {
	AppealID           int    `json:"appeal_id"`
	CitizenshipDetails string `json:"citizenship_details"`
	VoterCardDetails   string `json:"voter_card_details"`
	UserDetails        string `json:"user_details"`
	Status             string `json:"status"`
}

func CreateAppeal(appeal *Appeal) (int, error) {
	query := `INSERT INTO Appeals (CitizenshipDetails, VoterCardDetails, UserDetails, Status) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, appeal.CitizenshipDetails, appeal.VoterCardDetails, appeal.UserDetails, appeal.Status)
	if err != nil {
		return 0, err
	}

	var appealID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&appealID)
	if err != nil {
		return 0, err
	}

	return appealID, nil
}

func UpdateAppeal(appealID int, updatedAppeal *Appeal) error {
	query := `UPDATE Appeals SET CitizenshipDetails = ?, VoterCardDetails = ?, UserDetails = ?, Status = ? WHERE AppealID = ?`
	_, err := db.DB.Exec(query, updatedAppeal.CitizenshipDetails, updatedAppeal.VoterCardDetails, updatedAppeal.UserDetails, updatedAppeal.Status, appealID)
	return err
}

func GetAppeals(limit, offset int) ([]Appeal, error) {
	query := `SELECT AppealID, CitizenshipDetails, VoterCardDetails, UserDetails, Status FROM Appeals ORDER BY AppealID DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appeals []Appeal
	for rows.Next() {
		var appeal Appeal
		if err := rows.Scan(&appeal.AppealID, &appeal.CitizenshipDetails, &appeal.VoterCardDetails, &appeal.UserDetails, &appeal.Status); err != nil {
			return nil, err
		}
		appeals = append(appeals, appeal)
	}
	return appeals, nil
}

func GetAppealByID(appealID int) (*Appeal, error) {
	query := `SELECT AppealID, CitizenshipDetails, VoterCardDetails, UserDetails, Status FROM Appeals WHERE AppealID = ?`
	var appeal Appeal
	err := db.DB.QueryRow(query, appealID).Scan(&appeal.AppealID, &appeal.CitizenshipDetails, &appeal.VoterCardDetails, &appeal.UserDetails, &appeal.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Appeal With That ID")
		}
		return nil, err
	}
	return &appeal, nil
}

func DeleteAppeal(appealID int) error {
	query := `DELETE FROM Appeals WHERE AppealID = ?`
	_, err := db.DB.Exec(query, appealID)
	return err
}
