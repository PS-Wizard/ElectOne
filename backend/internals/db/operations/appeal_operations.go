package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type Appeal struct {
	AppealID      int    `json:"appeal_id"`
	CitizenshipID string `json:"citizenship_id"`
	VoterCardID   string `json:"voter_card_id"`
	Password      string `json:"password"`
	Photos        string `json:"photos"`
	Status        string `json:"status,omitempty"`
}

func CreateAppeal(appeal *Appeal) (int, error) {
	appeal.Status = "Pending"
	query := `INSERT INTO Appeals (CitizenshipID, VoterCardID, Password, Photos, Status) VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, appeal.CitizenshipID, appeal.VoterCardID, appeal.Password, appeal.Photos, appeal.Status)
	if err != nil {
		return 0, err
	}

	var appealID int
	err = db.DB.QueryRow("SELECT last_insert_rowid()").Scan(&appealID)
	return appealID, err
}

func UpdateAppeal(appealID int, updatedAppeal *Appeal) error {
	query := `UPDATE Appeals SET CitizenshipID = ?, VoterCardID = ?, Password = ?, Photos = ?, Status = ? WHERE AppealID = ?`
	_, err := db.DB.Exec(query, updatedAppeal.CitizenshipID, updatedAppeal.VoterCardID, updatedAppeal.Password, updatedAppeal.Photos, updatedAppeal.Status, appealID)
	return err
}

func GetAppeals(limit, offset int) ([]Appeal, error) {
	query := `SELECT AppealID, CitizenshipID, VoterCardID, Password, Photos, Status FROM Appeals ORDER BY AppealID DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var appeals []Appeal
	for rows.Next() {
		var a Appeal
		if err := rows.Scan(&a.AppealID, &a.CitizenshipID, &a.VoterCardID, &a.Password, &a.Photos, &a.Status); err != nil {
			return nil, err
		}
		appeals = append(appeals, a)
	}
	return appeals, nil
}

func GetAppealByID(appealID int) (*Appeal, error) {
	query := `SELECT AppealID, CitizenshipID, VoterCardID, Password, Photos, Status FROM Appeals WHERE AppealID = ?`
	var a Appeal
	err := db.DB.QueryRow(query, appealID).Scan(&a.AppealID, &a.CitizenshipID, &a.VoterCardID, &a.Password, &a.Photos, &a.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Appeal With That ID")
		}
		return nil, err
	}
	return &a, nil
}

func GetAppealsCount() (int, error) {
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM Appeals").Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func ApproveAppeal(appealID int) error {
	appeal, err := GetAppealByID(appealID)
	if err != nil {
		return err
	}

	// Create a new user from the approved appeal
	user := &User{
		CitizenshipID: appeal.CitizenshipID,
		VoterCardID:   appeal.VoterCardID,
		Password:      appeal.Password,
		Photos:        appeal.Photos,
		TotpSecret:    "",
	}

	_, err = CreateUser(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	appeal.Status = "approved"
	return UpdateAppeal(appealID, appeal)
}

func RejectAppeal(appealID int) error {
	appeal, err := GetAppealByID(appealID)
	if err != nil {
		return err
	}

	appeal.Status = "rejected"
	return UpdateAppeal(appealID, appeal)
}

func DeleteAppeal(id int) error {
	return RejectAppeal(id)
}
