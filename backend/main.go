package main

import (
	"fmt"
	"net/http"

	"github.com/PS-Wizard/VotingSystem/models"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /newuser", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid Form", http.StatusBadRequest)
			return
		}

		id := r.FormValue("voterid")
		password := r.FormValue("password")
		models.CreateUser(id, password)
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
