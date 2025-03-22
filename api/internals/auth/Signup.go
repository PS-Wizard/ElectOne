package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PS-Wizard/VotingSystem/db"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	CitizenID string `json:"citizenid"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

func (sr *SignupRequest) NewUser() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(sr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	query := `INSERT INTO users (citizenid, password, phone) VALUES (?, ?, ?)`
	_, err = db.DB.Exec(query, sr.CitizenID, hashedPassword, sr.Phone)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("user already exists")
		}
		return fmt.Errorf("failed to insert user into database: %v", err)
	}

	fmt.Println("User Successfully Created!")
	return nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	var signUpReq SignupRequest
	err := json.NewDecoder(r.Body).Decode(&signUpReq)
	if err != nil {
		jsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = signUpReq.NewUser()
	if err != nil {
		jsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Send success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Successfully Created"})
}
