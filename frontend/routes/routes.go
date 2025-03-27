package routes

import (
	"github.com/PS-Wizard/ElectOneui/views/pages/home"
	"github.com/gofiber/fiber/v2"
)

func HandleRoutes(ctx *fiber.App) error {
	ctx.Get("/", home.RenderHome)
	return nil
}
