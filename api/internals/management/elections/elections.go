package elections

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PS-Wizard/VotingSystem/db"
	"github.com/PS-Wizard/VotingSystem/utils"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Election struct {
	ElectionID      uint   `json:"election_id"`
	Name            string `json:"name"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	VoteRestriction string `json:"vote_restriction"`
}

func (er *Election) createElection() error {
	query := `INSERT INTO elections (name, start_date, end_date, vote_restriction) VALUES (?, ?, ?, ?)`
	_, err := db.DB.Exec(query, er.ElectionID, er.Name, er.StartDate, er.EndDate, er.VoteRestriction)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("election with this ID already exists")
		}
		return fmt.Errorf("failed to insert election into database: %v", err)
	}

	fmt.Println("Election Successfully Created!")
	return nil
}

func CreateElectionHandler(w http.ResponseWriter, r *http.Request) {
	var electionReq Election

	err := json.NewDecoder(r.Body).Decode(&electionReq)
	if err != nil {
		utils.JsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = electionReq.createElection()
	if err != nil {
		utils.JsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Election Successfully Created"})
}

func getElections() ([]Election, error) {
	var elections []Election

	query := `SELECT election_id, name, start_date, end_date, vote_restriction FROM elections`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Error fetching elections: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var election Election
		err := rows.Scan(&election.ElectionID, &election.Name, &election.StartDate, &election.EndDate, &election.VoteRestriction)
		if err != nil {
			return nil, fmt.Errorf("Error scanning election: %v", err)
		}
		elections = append(elections, election)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error during row iteration: %v", err)
	}

	return elections, nil
}

func GetElectionsHandler(w http.ResponseWriter, r *http.Request) {
	elections, err := getElections()
	if err != nil {
		utils.JsonError(w, fmt.Sprintf("Error fetching elections: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(elections); err != nil {
		utils.JsonError(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}
