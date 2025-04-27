package routes

import (
	"log"

	"github.com/PS-Wizard/ElectOneAPI/api"
	appeals "github.com/PS-Wizard/ElectOneAPI/api/Appeals"
	auth "github.com/PS-Wizard/ElectOneAPI/api/Auth"
	candidates "github.com/PS-Wizard/ElectOneAPI/api/Candidates"
	citizens "github.com/PS-Wizard/ElectOneAPI/api/Citizens"
	elections "github.com/PS-Wizard/ElectOneAPI/api/Elections"
	redis "github.com/PS-Wizard/ElectOneAPI/api/Redis"
	users "github.com/PS-Wizard/ElectOneAPI/api/Users"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {
	// Citizen Routes:
	app.Get("/api/secure/citizens/:id", TokenValidationBoth, citizens.HandleSearch)
	app.Get("/api/secure/citizensPaginated/:offset", TokenValidationAdmin, citizens.HandleGet)
	app.Post("/api/secure/citizens", TokenValidationAdmin, citizens.HandleCreate)
	app.Put("/api/secure/citizens/:id", TokenValidationAdmin, citizens.HandleUpdate)
	app.Delete("/api/secure/citizens/:id", TokenValidationAdmin, citizens.HandleDelete)

	// User Routes
	app.Get("/api/secure/user/:id", TokenValidationBoth, users.HandleGetUser)
	app.Get("/api/secure/usersPaginated/:offset", TokenValidationAdmin, users.HandleGetUsersPaginated)
	app.Post("/api/secure/user", TokenValidationAdmin, users.HandleCreateNewUser)
	app.Put("/api/secure/user/:id", TokenValidationAdmin, users.HandleUpdateUserDetails)
	app.Delete("/api/secure/user/:id", TokenValidationAdmin, users.HandleDeleteUser)

	// Election Routes
	app.Get("/api/secure/election/:id", TokenValidationBoth, elections.HandleGetElection)
	app.Get("/api/secure/electionsPaginated/:offset", TokenValidationBoth, elections.HandleGetElectionsPaginated)
	app.Post("/api/secure/election", TokenValidationAdmin, elections.HandleCreateNewElection)
	app.Put("/api/secure/election/:id", TokenValidationAdmin, elections.HandleUpdateElectionDetails)
	app.Delete("/api/secure/election/:id", TokenValidationAdmin, elections.HandleDeleteElection)

	//Candidate Routes
	app.Get("/api/secure/candidate/:id", TokenValidationBoth, candidates.HandleGetCandidate)
	app.Get("/api/secure/candidatesPaginated/:offset", TokenValidationBoth, candidates.HandleGetCandidatesPaginated)
	app.Post("/api/secure/candidate", TokenValidationAdmin, candidates.HandleCreateCandidate)
	app.Put("/api/secure/candidate/:id", TokenValidationAdmin, candidates.HandleUpdateCandidate)
	app.Delete("/api/secure/candidate/:id", TokenValidationAdmin, candidates.HandleDeleteCandidate)

	app.Get("/api/secure/appeals/", TokenValidationAdmin, appeals.HandleGetRegistrationAppeals)
	app.Post("/api/secure/appeals/:id", TokenValidationAdmin, appeals.HandleAppealApprove)

	// Admin Login Routes:
	app.Post("/api/admin/signup", auth.HandleCreateAdmin)
	app.Post("/api/admin/login", auth.HandleAdminLogin)

	// User Login
	app.Post("/api/userlogin", auth.HandleUserLogin)
	app.Get("/api/secure/qr/:otpURL", auth.HandleGenerateQRCode)
	app.Get("/api/secure/verifyOTP", auth.ValidateTOTP)

	// cast vote test
	app.Post("/api/castVote", redis.HandleVoteIncrement)
	// Cast Vote:
	// app.Post("/api/castVote", TokenValidationUser, redis.HandleVoteIncrement)
	// Server Sent Event For Live Vote Count
	app.Get("/api/voteStream", websocket.New(func(c *websocket.Conn) {
		// Subscribe to Redis channel
		pubsub := api.RDB.PSubscribe(api.CTX, "voteUpdates")
		defer pubsub.Close()

		// Listen for Redis messages and send them to the WebSocket client
		for {
			msg, err := pubsub.ReceiveMessage(api.CTX)
			if err != nil {
				log.Println("Error receiving Redis message:", err)
				break
			}

			// Send the message to the WebSocket client
			err = c.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			if err != nil {
				log.Println("Error sending WebSocket message:", err)
				break
			}
		}
	}))

}
