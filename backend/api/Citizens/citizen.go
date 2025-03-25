package citizens

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"

	"github.com/PS-Wizard/ElectOneAPI/api"
)

func getCitizenById(citizenID string, db *sql.DB) (*citizen, error) {
	var citizen citizen
	query := "SELECT citizenID, fullName, dateOfBirth, placeOfResidence FROM citizens WHERE citizenID = ?"
	row := db.QueryRow(query, citizenID)
	err := row.Scan(&citizen.CitizenID, &citizen.Name, &citizen.DOB, &citizen.POR)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("citizen not found")
		}

		return nil, fmt.Errorf("failed to fetch citizen: %v", err)
	}
	return &citizen, nil
}

func getCitizensPaginated(offset int) ([]citizen, error) {
	var citizens []citizen
	query := "SELECT citizenID, fullName, dateOfBirth, placeOfResidence FROM citizens LIMIT ? OFFSET ?"
	rows, err := api.DB.Query(query, 10, offset) // The 10 here is the limit of citizens to fetch each time .
	if err != nil {
		return nil, fmt.Errorf("failed to fetch citizens: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var citizen citizen
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
