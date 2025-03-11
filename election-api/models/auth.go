package models

import (
	"regexp"
)

var JwtKey = []byte("superSecretKey")

type LoginRequest struct {
	VoterID  string `json:"voterid"`
	Password string `json:"password"`
}

// TODO: Need More Validation
type SignUpRequest struct {
	VoterID  string `json:"voterid"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (s *SignUpRequest) Validate() (string, bool) {
	/* TODO:
	   - Frontend: Needs to make sure password is a secure password; uppercase lowercase yada yada
	   - Frontend: Needs to make sure VoterID is a seemingly valid VoterID; business analyst needs to find out what a voterID looks like.
	*/
	match, _ := regexp.MatchString(`^\d{6}$`, s.VoterID)
	if !match {
		return "VoterID Doesn't Match Format", false
	}
	if len(s.Phone) < 10 {
		return "Invalid Phone Number Length", false
	}
	// Ensure password is not empty
	if len(s.Password) < 6 {
		return "Invalid Password Length", false
	}
	return "All Valid", true
}

func (lr *LoginRequest) Validate() (string, bool) {
	/* TODO:
	   - Frontend: Needs to make sure password is a secure password; uppercase lowercase yada yada
	   - Frontend: Needs to make sure VoterID is a seemingly valid VoterID; business analyst needs to find out what a voterID looks like.
	*/
	match, _ := regexp.MatchString(`^\d{6}$`, lr.VoterID)
	if !match {
		return "VoterID Doesn't Match Format", false
	}
	// Ensure password is not empty
	if len(lr.Password) < 6 {
		return "Invalid Password Length", false
	}
	return "All Valid", true
}
