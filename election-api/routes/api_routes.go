package routes

import (
	"net/http"

	"github.com/PS-Wizard/VotingSystem-API/controllers"
)

func HandleAPIRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/signup", controllers.HandleSignup)
	mux.HandleFunc("/api/login", controllers.HandleLogin)
}
