package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/PS-Wizard/VotingSystem/db"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type LoginRequest struct {
	CitizenID string `json:"citizenid"`
	Password  string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"jwt"`
}

var jwtSecret = []byte("supersecret69420")

func (lr *LoginRequest) GenerateJWT() (string, error) {
	claims := jwt.MapClaims{
		"citizenid": lr.CitizenID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)

}

func (lr *LoginRequest) LoginUser() error {
	var hashedPassword string
	query := "SELECT password FROM users WHERE citizenid = ?"
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

func jsonError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		jsonError(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = loginRequest.LoginUser()
	if err != nil {
		jsonError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenResponse, err := loginRequest.GenerateJWT()
	if err != nil {
		jsonError(w, "Can't Produce JWT", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: tokenResponse})
}
