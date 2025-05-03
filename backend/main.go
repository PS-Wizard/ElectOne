package main

import (
	"log"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/PS-Wizard/Electone/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitTurso()
	db.InitRedis()
	defer db.CloseTurso()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

