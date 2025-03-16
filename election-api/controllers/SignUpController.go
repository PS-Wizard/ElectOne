package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PS-Wizard/VotingSystem-API/models"
	"golang.org/x/crypto/bcrypt"
)

var db = []models.SignUpRequest{}
func HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req models.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		log.Printf("Failed Decoding: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request"))
		return
	}

	if resp, err := req.Validate(); !err {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(resp))
        return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed Hashing The Password: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error #0"))
		return
	}
	db = append(db, models.SignUpRequest{
		VoterID:  req.VoterID,
		Phone:    req.Phone,
		Password: string(hashedPassword),
	})
}
