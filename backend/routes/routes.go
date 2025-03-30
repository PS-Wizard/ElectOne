package routes

import (
	auth "github.com/PS-Wizard/ElectOneAPI/api/Auth"
	candidates "github.com/PS-Wizard/ElectOneAPI/api/Candidates"
	citizens "github.com/PS-Wizard/ElectOneAPI/api/Citizens"
	elections "github.com/PS-Wizard/ElectOneAPI/api/Elections"
	redis "github.com/PS-Wizard/ElectOneAPI/api/Redis"
	users "github.com/PS-Wizard/ElectOneAPI/api/Users"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {
	// Citizen Routes:
	app.Get("/api/secure/citizens/:id", citizens.HandleSearch)
	app.Get("/api/secure/citizensPaginated/:offset", citizens.HandleGet)
	app.Post("/api/secure/citizens", citizens.HandleCreate)
	app.Put("/api/secure/citizens/:id", citizens.HandleUpdate)
	app.Delete("/api/secure/citizens/:id", citizens.HandleDelete)

	// User Routes
	app.Get("/api/secure/user/:id", users.HandleGetUser)
	app.Get("/api/secure/usersPaginated/:offset", users.HandleGetUsersPaginated)
	app.Post("/api/secure/user", users.HandleCreateNewUser)
	app.Put("/api/secure/user/:id", users.HandleUpdateUserDetails)
	app.Delete("/api/secure/user/:id", users.HandleDeleteUser)

	// Election Routes
	app.Get("/api/secure/election/:id", elections.HandleGetElection)
	app.Get("/api/secure/electionsPaginated/:offset", elections.HandleGetElectionsPaginated)
	app.Post("/api/secure/election", elections.HandleCreateNewElection)
	app.Put("/api/secure/election/:id", elections.HandleUpdateElectionDetails)
	app.Delete("/api/secure/election/:id", elections.HandleDeleteElection)

	//Candidate Routes
	app.Get("/api/secure/candidate/:id", candidates.HandleGetCandidate)
	app.Get("/api/secure/candidatesPaginated/:offset", candidates.HandleGetCandidatesPaginated)
	app.Post("/api/secure/candidate", candidates.HandleCreateCandidate)
	app.Put("/api/secure/candidate/:id", candidates.HandleUpdateCandidate)
	app.Delete("/api/secure/candidate/:id", candidates.HandleDeleteCandidate)

	// Admin Login Routes:
	app.Post("/api/admin/signup", auth.HandleCreateAdmin)
	app.Post("/api/admin/login", auth.HandleAdminLogin)

	app.Post("/api/userlogin", auth.HandleUserLogin)

	// Cast Vote:
	app.Post("/api/castVote", redis.HandleVoteIncrement)

}
