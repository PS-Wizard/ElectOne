package models

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	VoterID  string `json:"voterID"`
	Password string
}

var ListOfUsers = []User{}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func CreateUser(id, password string) error {
	hashedPassword, err := HashPassword(password)
	fmt.Println("Saved Hash Successfully :%s", hashedPassword)
	if err != nil {
		return err
	}
	ListOfUsers = append(ListOfUsers, User{VoterID: id, Password: hashedPassword})
	return nil
}

func CheckCredentials(id, password string) bool {
	for _, user := range ListOfUsers {
		if user.VoterID == id {
			return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
		}
	}
	return false
}

func GenerateSessionToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid Form", http.StatusBadRequest)
		return
	}

	id := r.FormValue("voterid")
	password := r.FormValue("password")

	if CheckCredentials(id, password) {
		token := GenerateSessionToken()
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: token,
			Path:  "/",
		})
		fmt.Fprintf(w, "Login Successful")
	} else {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
	}
}
