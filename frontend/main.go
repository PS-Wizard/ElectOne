package main

import (
	"log"

	"github.com/PS-Wizard/ElectOneui/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./static")
	routes.HandleRoutes(app)
	log.Fatal(app.Listen(":5173"))
}
