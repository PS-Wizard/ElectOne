package routes

import (
	"github.com/PS-Wizard/Electone/internals/auth"
	"github.com/PS-Wizard/Electone/internals/db/handlers"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/PS-Wizard/Electone/realtime"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	// Auth Routes ( DOESNT REQUIRE JWT)
	authRoutes := app.Group("/auth")
	app.Get("/vote/ws", websocket.New(realtime.HandleLivePushUpdates))
	authRoutes.Post("/user", auth.UserLogin)
	authRoutes.Post("/admin", auth.AdminLogin)

	protected := app.Group("/", middlewares.JWTMiddleware())
	// Citizenship routes
	citizenshipRoutes := protected.Group("/citizen")
	handlers.RegisterCitizenshipRoutes(citizenshipRoutes)

	// VoterCard Routes
	voterCardRoutes := protected.Group("/voter_card")
	handlers.RegisterVoterCardRoutes(voterCardRoutes)

	// User Routes
	userRoutes := protected.Group("/user")
	handlers.RegisterUserRoutes(userRoutes)

	// Election Routes
	electionRoutes := protected.Group("/election")
	handlers.RegisterElectionRoutes(electionRoutes)

	// Appeal Routes
	appealRoutes := protected.Group("/appeal")
	handlers.RegisterAppealRoutes(appealRoutes)

	// Candidate Routes
	candidateRoutes := protected.Group("/candidate")
	handlers.RegisterCandidateRoutes(candidateRoutes)

	// Vote Routes

	protected.Post("/vote", middlewares.RequireUser, realtime.HandleVote)

	//admin Routes
	adminRoutes := protected.Group("/admin")
	handlers.RegisterAdminRoutes(adminRoutes)
}
