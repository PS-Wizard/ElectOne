package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PS-Wizard/VotingSystem/db"
	"github.com/PS-Wizard/VotingSystem/utils"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	CitizenID string `json:"citizen_id"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"jwt"`
}

var jwtSecret = []byte("supersecret69420")

func (lr *LoginRequest) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"citizen_id": lr.CitizenID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func (lr *LoginRequest) LoginUser() error {
	var hashedPassword string
	query := "SELECT password FROM users WHERE citizen_id = ?"
	err := db.DB.QueryRow(query, lr.CitizenID).Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found")
		}
		return fmt.Errorf("failed to query user: %v\n", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(lr.Password))
	if err != nil {
		return fmt.Errorf("Incorrect Password")
	}

	return nil
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.JsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = loginRequest.LoginUser()
	if err != nil {
		utils.JsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenResponse, err := loginRequest.GenerateJWT()
	if err != nil {
		utils.JsonError(w, "Can't Produce JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenResponse})
}
