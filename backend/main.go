package main

import (
	"log"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/PS-Wizard/Electone/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		// AllowOrigins: "http://localhost:5173",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	db.InitTurso()
	db.InitRedis()
	defer db.CloseTurso()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
