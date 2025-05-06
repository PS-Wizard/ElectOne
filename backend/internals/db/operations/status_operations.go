package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type AppealStatusReq struct {
	AppealID      int    `json:"appeal_id"`
	CitizenshipID string `json:"citizenship_id"`
	VoterCardID   string `json:"voter_card_id"`
	Password      string `json:"password"`
	Status        string `json:"status,omitempty"`
}

func GetAppealStatus(request *AppealStatusReq) (*Appeal, error) {

	query := `
		SELECT AppealID, CitizenshipID, VoterCardID, Password, Photos, Status 
		FROM Appeals 
		WHERE AppealID = ? AND CitizenshipID = ? AND VoterCardID = ? AND Password = ?`

	var a Appeal

	err := db.DB.QueryRow(query, request.AppealID, request.CitizenshipID, request.VoterCardID, request.Password).
		Scan(&a.AppealID, &a.CitizenshipID, &a.VoterCardID, &a.Password, &a.Photos, &a.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no appeal found with the provided details")
		}
		return nil, err
	}

	// Return the appeal if a match is found
	return &a, nil
}
