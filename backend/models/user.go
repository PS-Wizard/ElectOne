package models

import (
	"fmt"
)

type User struct {
	VoterID  string `json:"voterID"`
	password string
}

// TODO: Hash The Password
var ListOfUsers = []User{
	{
		VoterID:  "123456",
		password: "SomePassword",
	},
}

func CreateUser(id, password string) {
	fmt.Printf("Got %s, %s \n", id, password)
	ListOfUsers = append(ListOfUsers, User{VoterID: id, password: password})
}
