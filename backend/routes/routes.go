package routes

import (
	citizens "github.com/PS-Wizard/ElectOneAPI/api/Citizens"
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

}
