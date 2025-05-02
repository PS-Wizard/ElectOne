package operations

import (
	"database/sql"
	"fmt"

	"github.com/PS-Wizard/Electone/internals/db"
)

type Citizenship struct {
	CitizenID     string `json:"citizenship_id"`
	FullName      string `json:"full_name"`
	DateOfBirth   string `json:"date_of_birth"`
	BirthPlace    string `json:"birth_place"`
	PermanentAddr string `json:"permanent_address"`
}

func CreateCitizenship(citizenship *Citizenship) (string, error) {
	query := `INSERT INTO Citizenship (CitizenID, FullName, DateOfBirth, BirthPlace, PermanentAddress) VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, citizenship.CitizenID, citizenship.FullName, citizenship.DateOfBirth, citizenship.BirthPlace, citizenship.PermanentAddr)
	if err != nil {
		return "", err
	}

	return citizenship.CitizenID, nil
}

func UpdateCitizenship(citizenID string, updatedCitizenship *Citizenship) error {
	query := `UPDATE Citizenship SET FullName = ?, DateOfBirth = ?, BirthPlace = ?, PermanentAddress = ? WHERE CitizenID = ?`
	_, err := db.DB.Exec(query, updatedCitizenship.FullName, updatedCitizenship.DateOfBirth, updatedCitizenship.BirthPlace, updatedCitizenship.PermanentAddr, citizenID)
	return err
}

func GetCitizenships(limit, offset int) ([]Citizenship, error) {
	query := `SELECT CitizenID, FullName, DateOfBirth, BirthPlace, PermanentAddress FROM Citizenship ORDER BY CitizenID DESC LIMIT ? OFFSET ?`
	rows, err := db.DB.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var citizenships []Citizenship
	for rows.Next() {
		var citizenship Citizenship
		if err := rows.Scan(&citizenship.CitizenID, &citizenship.FullName, &citizenship.DateOfBirth, &citizenship.BirthPlace, &citizenship.PermanentAddr); err != nil {
			return nil, err
		}
		citizenships = append(citizenships, citizenship)
	}
	return citizenships, nil
}

func GetCitizenshipByID(citizenID string) (*Citizenship, error) {
	query := `SELECT CitizenID, FullName, DateOfBirth, BirthPlace, PermanentAddress FROM Citizenship WHERE CitizenID = ?`
	var citizenship Citizenship
	err := db.DB.QueryRow(query, citizenID).Scan(&citizenship.CitizenID, &citizenship.FullName, &citizenship.DateOfBirth, &citizenship.BirthPlace, &citizenship.PermanentAddr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("No Citizen With THat ID")
		}
		return nil, err
	}
	return &citizenship, nil
}

func DeleteCitizenship(db *sql.DB, citizenID string) error {
	query := `DELETE FROM Citizenship WHERE CitizenID = ?`
	_, err := db.Exec(query, citizenID)
	return err
}
