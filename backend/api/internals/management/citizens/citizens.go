package citizens

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PS-Wizard/VotingSystem/db"
	"github.com/PS-Wizard/VotingSystem/utils"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Citizen struct {
	CitizenID        string `json:"citizen_id"`
	FullName         string `json:"full_name"`
	DOB              string `json:"dob"`
	POB              string `json:"pob"`
	PermanentAddress string `json:"permanent_address"`
}

func (cr *Citizen) CreateCitizen() error {
	query := `INSERT INTO citizens (citizen_id, full_name, dob, pob, permanent_address) VALUES (?, ?, ?, ?, ?)`
	_, err := db.DB.Exec(query, cr.CitizenID, cr.FullName, cr.DOB, cr.POB, cr.PermanentAddress)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("citizen with this ID already exists")
		}
		return fmt.Errorf("failed to insert citizen into database: %v", err)
	}

	fmt.Println("Citizen Successfully Created!")
	return nil
}

func CreateCitizenHandler(w http.ResponseWriter, r *http.Request) {
	var citizenReq Citizen

	err := json.NewDecoder(r.Body).Decode(&citizenReq)
	if err != nil {
		utils.JsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = citizenReq.CreateCitizen()
	if err != nil {
		utils.JsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Citizen Successfully Created"})
}
