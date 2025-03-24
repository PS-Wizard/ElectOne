package routes

import (
	"github.com/PS-Wizard/VotingSystemUI/views"
	"github.com/PS-Wizard/VotingSystemUI/views/pages"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		homePage := pages.Home()
		return views.Render(c, homePage)

	})

}
