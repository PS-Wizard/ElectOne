package routes

import (
	"time"

	"github.com/PS-Wizard/Electone/internals/auth"
	"github.com/PS-Wizard/Electone/internals/db/handlers"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/PS-Wizard/Electone/statistics"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App) {

	rateLimiter := limiter.New(limiter.Config{
		Max:               5,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	})

	// Auth Routes ( DOESNT REQUIRE JWT)
	authRoutes := app.Group("/auth")
	app.Get("/vote/ws", websocket.New(statistics.HandleLivePushUpdates))
	app.Get("/vote/history/:electionID", statistics.GetElectionHistory)

	authRoutes.Post("/user", rateLimiter, middlewares.RequireLocation, auth.UserLogin)
	authRoutes.Post("/admin", rateLimiter, auth.AdminLogin)

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
	// TODO: MIGHT WANT TO DISABLE CACHE
	electionRoutes := protected.Group("/election", cache.New())
	// electionRoutes := protected.Group("/election")
	handlers.RegisterElectionRoutes(electionRoutes)

	// Appeal Routes
	appealRoutes := protected.Group("/appeal")
	handlers.RegisterAppealRoutes(appealRoutes)

	// Candidate Routes
	candidateRoutes := protected.Group("/candidate")
	handlers.RegisterCandidateRoutes(candidateRoutes)

	// Vote Routes

	protected.Post("/vote", middlewares.RequireUser, statistics.HandleVote)

	//admin Routes
	adminRoutes := protected.Group("/admin")
	handlers.RegisterAdminRoutes(adminRoutes)
}
