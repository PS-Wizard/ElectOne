package candidates

import (
	"encoding/json"
	"net/http"

	"github.com/PS-Wizard/VotingSystem/utils"
)

func CreateRunnerHandler(w http.ResponseWriter, r *http.Request) {
	var runnerReq Candidate
	err := json.NewDecoder(r.Body).Decode(&runnerReq)
	if err != nil {
		utils.JsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = runnerReq.CreateRunner()
	if err != nil {
		utils.JsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Runner Successfully Created"})
}

func GetRunnersHandler(w http.ResponseWriter, r *http.Request) {
	runners, err := GetRunners()
	if err != nil {
		utils.JsonError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(runners)
}
