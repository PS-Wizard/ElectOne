package routes

import (
	"github.com/PS-Wizard/VotingSystem/internals/auth"
	"net/http"
)

func HandleRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/login", auth.LoginHandler)
	mux.HandleFunc("POST /api/signup", auth.SignupHandler)
}
