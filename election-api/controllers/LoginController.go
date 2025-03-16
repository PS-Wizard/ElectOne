package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/PS-Wizard/VotingSystem-API/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Printf("Failed Decoding: %v\n", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if status, err := req.Validate(); !err {
		http.Error(w, status, http.StatusBadRequest)
		return
	}

	//TODO: Replace db with the actual database later
	for _, v := range db {
		if v.VoterID == req.VoterID {

			// Compare hashed password
			err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(req.Password))
			if err != nil {
				http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
				return
			}

			// Generate JWT token
			token, err := generateJWT(req.VoterID)
			if err != nil {
				log.Printf("JWT Error: %v\n", err)
				http.Error(w, "Internal Server Error #1", http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "token",
				Value:    token,
				Path:     "/",
				Expires:  time.Now().Add(24 * time.Hour),
				Secure:   false, // TODO: set to  TRUE when deploying
				HttpOnly: true,
			})

			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}
	}
	http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	http.Redirect(w, r, "/login?error=invalid", http.StatusUnauthorized)
}

func generateJWT(voterID string) (string, error) {
	claims := &jwt.MapClaims{
		"voterid": voterID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(models.JwtKey)
}
