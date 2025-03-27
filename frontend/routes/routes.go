package routes

import "github.com/gofiber/fiber/v2"

func HandleRoutes(ctx *fiber.App) error {
	ctx.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Write([]byte("Hello There"))
		return nil
	})
	return nil
}
