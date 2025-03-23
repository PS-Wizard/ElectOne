package routes

import (
	"net/http"

	"github.com/PS-Wizard/VotingSystem/internals/auth"
	"github.com/PS-Wizard/VotingSystem/internals/management/candidates"
	"github.com/PS-Wizard/VotingSystem/internals/management/citizens"
	"github.com/PS-Wizard/VotingSystem/internals/management/elections"
)

func HandleRoutes(mux *http.ServeMux) {
	// Auth Routes:
	mux.HandleFunc("POST /api/login", auth.LoginHandler)
	mux.HandleFunc("POST /api/signup", auth.SignupHandler)

	// Election Routes:
	mux.HandleFunc("GET /api/getElections", elections.GetElectionsHandler)
	mux.HandleFunc("POST /api/createElection", elections.CreateElectionHandler)

	// Create Citizen:
	mux.HandleFunc("POST /api/createCitizen", citizens.CreateCitizenHandler)

	// Runner Routes:
	mux.HandleFunc("POST /api/createCandidate", candidates.CreateRunnerHandler)
	mux.HandleFunc("GET /api/getCandidates", candidates.GetRunnersHandler)
}
