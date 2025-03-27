package home

import (
	"github.com/gofiber/fiber/v2"
)

func RenderHome(ctx *fiber.Ctx) error {
	ctx.Set("Content-type", "text/html")
	return home().Render(ctx.Context(), ctx.Response().BodyWriter())
}
