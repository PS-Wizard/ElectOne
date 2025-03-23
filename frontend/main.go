package main

import (
	"log"

	"github.com/PS-Wizard/VotingSystemUI/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Static("/static", "./views/static/")
	routes.HandleRoutes(app)
	log.Println("Server Up On :3000")
	log.Fatal(app.Listen(":3000"))

}
