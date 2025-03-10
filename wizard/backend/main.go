package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/PS-Wizard/VotingSystem/models"
	"github.com/PS-Wizard/VotingSystem/try"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./try/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "No session found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Error retrieving cookie", http.StatusBadRequest)
			return
		}

		fmt.Printf("Found Session Cookie: %s", cookie)
		err = try.Home().Render(context.Background(), w)
	})
	mux.Handle("/login", templ.Handler(try.LoginForm()))
	mux.Handle("/signup", templ.Handler(try.SignupForm()))

	mux.HandleFunc("POST /auth/newuser", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid Form", http.StatusBadRequest)
			return
		}

		id := r.FormValue("voterid")
		password := r.FormValue("password")

		if err := models.CreateUser(id, password); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created successfully")
	})

	mux.HandleFunc("POST /auth/login", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid Form", http.StatusBadRequest)
			return
		}

		id := r.FormValue("voterid")
		password := r.FormValue("password")

		if models.CheckCredentials(id, password) {
			token := models.GenerateSessionToken()
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    token,
				Path:     "/",
				HttpOnly: true,  // Prevents JS access
				Secure:   false, // Ensures it's only sent over HTTPS // NOTE: ENSURE THIS FOR PRODUCTION
			})
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Login Successful")
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		}
	})
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", mux)
}
