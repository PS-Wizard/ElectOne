package routes

import (
	citizens "github.com/PS-Wizard/ElectOneAPI/api/Citizens"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {
	// Citizen Routes:
	app.Get("/api/secure/citizens/:id", citizens.HandleSearch)
    app.Get("/api/secure/citizensPaginated/:offset", citizens.HandleGet)
	app.Post("/api/secure/citizens", citizens.HandleCreate)
	app.Put("/api/secure/citizens/:id", citizens.HandleUpdate)
	app.Delete("/api/secure/citizens/:id", citizens.HandleDelete)

}
