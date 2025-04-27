package appeals

import (
	"fmt"
	"strings"

	"github.com/PS-Wizard/ElectOneAPI/api"
	citizens "github.com/PS-Wizard/ElectOneAPI/api/Citizens"
	users "github.com/PS-Wizard/ElectOneAPI/api/Users"
)

// fetch pending registration appeals
func getRegistrationAppeals() ([]Appeal, error) {
	var appeals []Appeal
	query := "SELECT appealID, citizenID, fullName, dateOfBirth, placeOfResidence, photos, password, phoneNumber, tags FROM registration_appeals WHERE status = 'pending'"
	rows, err := api.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var appeal Appeal
		var photos string // assuming photos are stored as a string (maybe JSON or CSV)

		err := rows.Scan(&appeal.AppealID, &appeal.CitizenID, &appeal.Name, &appeal.DOB, &appeal.POR, &photos, &appeal.Password, &appeal.Phone, &appeal.Tag)
		if err != nil {
			return nil, err
		}

		// Split photos into a slice if they are stored as comma-separated
		if photos != "" {
			appeal.Photos = strings.Split(photos, ",")
		}

		appeals = append(appeals, appeal)
	}
	return appeals, nil
}

func approveRegistrationAppeal(appealID string) error {
	// Step 1: Fetch the appeal to get citizen and user details
	appeal, err := getAppealByID(appealID)
	if err != nil {
		return fmt.Errorf("failed to fetch appeal: %v", err)
	}

	// Step 2: Insert Citizen data
	err = citizens.CreateNewCitizen(citizens.Citizen{
		CitizenID: appeal.CitizenID,
		Name:      appeal.Name,
		DOB:       appeal.DOB,
		POR:       appeal.POR,
	})
	if err != nil {
		return fmt.Errorf("failed to insert citizen data: %v", err)
	}

	// Step 3: Insert User data
	err = users.CreateNewUser(users.User{
		CitizenID: appeal.CitizenID,
		Password:  appeal.Password,
		Phone:     appeal.Phone,
		Tag:       appeal.Tag,
		OTP:       "",                               // Empty OTP for now
		Photos:    strings.Join(appeal.Photos, ","), // Joining photos into a CSV string if they are an array
	})

	if err != nil {
		return fmt.Errorf("failed to insert user data: %v", err)
	}

	// Step 4: Update appeal status to 'approved'
	query := "UPDATE registration_appeals SET status = 'approved' WHERE appealID = ?"
	_, err = api.DB.Exec(query, appealID)
	if err != nil {
		return fmt.Errorf("failed to update appeal status: %v", err)
	}

	return nil
}

func getAppealByID(appealID string) (*Appeal, error) {
	var appeal Appeal
	query := "SELECT appealID, citizenID, fullName, dateOfBirth, placeOfResidence, photos, password, phoneNumber, tags FROM registration_appeals WHERE appealID = ?"
	row := api.DB.QueryRow(query, appealID)

	var photos string
	err := row.Scan(&appeal.AppealID, &appeal.CitizenID, &appeal.Name, &appeal.DOB, &appeal.POR, &photos, &appeal.Password, &appeal.Phone, &appeal.Tag)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch appeal: %v", err)
	}

	// Handle photos (comma-separated in DB)
	if photos != "" {
		appeal.Photos = strings.Split(photos, ",")
	}

	return &appeal, nil
}
