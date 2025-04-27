package citizens

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

func getCitizenById(citizenID string) (*Citizen, error) {
	var citizen Citizen
	query := "SELECT citizenID, fullName, dateOfBirth, placeOfResidence FROM citizens WHERE citizenID = ?"
	row := api.DB.QueryRow(query, citizenID)
	err := row.Scan(&citizen.CitizenID, &citizen.Name, &citizen.DOB, &citizen.POR)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("citizen not found")
		}

		return nil, fmt.Errorf("failed to fetch citizen: %v", err)
	}
	return &citizen, nil
}

func getCitizensPaginated(offset int) ([]Citizen, error) {
	var citizens []Citizen
	query := "SELECT citizenID, fullName, dateOfBirth, placeOfResidence FROM citizens LIMIT ? OFFSET ?"
	rows, err := api.DB.Query(query, 10, offset) // The 10 here is the limit of citizens to fetch each time .
	if err != nil {
		return nil, fmt.Errorf("failed to fetch citizens: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var citizen Citizen
		err := rows.Scan(&citizen.CitizenID, &citizen.Name, &citizen.DOB, &citizen.POR)
		if err != nil {
			return nil, fmt.Errorf("Failed To Scan Citizen: %v", err)
		}
		citizens = append(citizens, citizen)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}
	return citizens, nil

}

func CreateNewCitizen(citizen Citizen) error {
	query := "INSERT INTO citizens (citizenID, fullName, dateOfBirth, placeOfResidence) VALUES (?, ?, ?, ?)"
	_, err := api.DB.Exec(query, citizen.CitizenID, citizen.Name, citizen.DOB, citizen.POR)
	if err != nil {
		return fmt.Errorf("Failed to insert citizen details: %v", err)
	}
	return nil
}

func updateCitizenDetails(citizen Citizen, citizenID string) error {
	query := "UPDATE citizens SET fullName = ?, dateOfBirth = ?, placeOfResidence = ? WHERE citizenID = ?"
	_, err := api.DB.Exec(query, citizen.Name, citizen.DOB, citizen.POR, citizenID)
	if err != nil {
		return fmt.Errorf("Failed to update citizen details: %v", err)
	}
	return nil
}

func deleteCitizen(citizenID string) error {
	query := "DELETE FROM citizens WHERE citizenID = ?"
	_, err := api.DB.Exec(query, citizenID)
	if err != nil {
		return fmt.Errorf("Failed to delete citizen: %v", err)
	}
	return nil
}
