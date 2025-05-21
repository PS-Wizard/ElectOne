package routes

import (
	"time"

	"github.com/PS-Wizard/Electone/internals/auth"
	"github.com/PS-Wizard/Electone/internals/db/handlers"
	"github.com/PS-Wizard/Electone/internals/ws"

	// "github.com/PS-Wizard/Electone/internals/db/operations"
	"github.com/PS-Wizard/Electone/internals/middlewares"
	"github.com/gofiber/fiber/v2"

	// "github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func SetupRoutes(app *fiber.App) {

	rateLimiter := limiter.New(limiter.Config{
		Max:               5,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	})

	hub := ws.NewHub()
	ws.SetupWebSocketRoute(app.Group("/live"), hub)
	handlers.Hub = hub
	// Auth Routes ( DOESNT REQUIRE JWT)
	authRoutes := app.Group("/auth")

	// authRoutes.Post("/user", rateLimiter, middlewares.RequireLocation, auth.UserLogin)
	authRoutes.Post("/user", rateLimiter, auth.UserLogin)
	authRoutes.Post("/admin", rateLimiter, auth.AdminLogin)
	authRoutes.Post("/forgot", rateLimiter, auth.HandleForgotPassword)
	authRoutes.Post("/delete", rateLimiter, auth.HandleDeleteUser)

	// Appeal Status Routes (Doesnt Require JWT)
	statusAppealRoutes := app.Group("/status", rateLimiter)
	handlers.RegisterStatusRoutes(statusAppealRoutes)

	appealRoutes := app.Group("/appeal")
	appealRoutes.Post("/", handlers.CreateAppealHandler)
	appealRoutes.Get("/count", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.GetAppealsCountHandler)
	appealRoutes.Get("/:id", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.GetAppealByIDHandler)
	appealRoutes.Put("/:id", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.UpdateAppealHandler)
	appealRoutes.Delete("/:id", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.DeleteAppealHandler)
	appealRoutes.Get("/", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.ListAppealsHandler)

	appealRoutes.Post("/:id/approve", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.ApproveAppealHandler)
	appealRoutes.Post("/:id/reject", middlewares.JWTMiddleware(), middlewares.RequireAdmin, handlers.RejectAppealHandler)

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
	// electionRoutes := protected.Group("/election", cache.New())
	electionRoutes := protected.Group("/election")
	handlers.RegisterElectionRoutes(electionRoutes)

	// Appeal Routes

	// Candidate Routes
	candidateRoutes := protected.Group("/candidate")
	handlers.RegisterCandidateRoutes(candidateRoutes)

	// Turso Vote Routes
	// votingRoutes := protected.Group("/vote")
	// handlers.RegisterVoteRoutes(votingRoutes)

	tursoVotingRoutes := protected.Group("/vote/turso")
	handlers.RegisterVoteRoutes(tursoVotingRoutes)

	// Redis Vote Routes
	redisVotingRoutes := protected.Group("/vote/redis")
	handlers.RegisterRedisVotingRoutes(redisVotingRoutes)

	// Sync To Turso:
	syncRoute := protected.Group("/sync")
	handlers.RegisterSyncRoutes(syncRoute)

	//admin Routes
	adminRoutes := protected.Group("/admin")
	handlers.RegisterAdminRoutes(adminRoutes)

}
