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
	app.Static("/uploads", "./uploads/")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // or specific origin for more security
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "*",
	}))

	db.InitTurso()
	defer db.CloseTurso()
	db.InitRedis()
	defer db.CloseRedis()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
