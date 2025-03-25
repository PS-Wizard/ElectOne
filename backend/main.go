package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/PS-Wizard/ElectOneAPI/api"
	"github.com/PS-Wizard/ElectOneAPI/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	authToken := os.Getenv("TURSO_AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("Turso's Authentication Token Missing From Enviroment")
	}

	dbURL := os.Getenv("TURSO_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("Turso's Database URL Missing From Enviroment")
	}

	api.InitializeDB(authToken, dbURL)
	defer api.CloseDB()

	f, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed To Open Log File: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	app := fiber.New(fiber.Config{
		Prefork:                 false,                                // Use multiple OS threads for better performance
		ServerHeader:            "ElectOneAPI",                        // Custom server header for branding
		AppName:                 "ElectOneAPI",                        // Useful for debugging/logs
		StrictRouting:           false,                                // /route and /route/ are the same (easier UX)
		CaseSensitive:           false,                                // Routes are NOT case-sensitive (/User == /user)
		BodyLimit:               10 * 1024 * 1024,                     // 10MB max request body size
		ReadTimeout:             10 * time.Second,                     // Prevents slow attacks
		WriteTimeout:            10 * time.Second,                     // Prevents response delays
		IdleTimeout:             30 * time.Second,                     // Closes idle connections
		EnablePrintRoutes:       true,                                 // Print all routes on startup
		DisableStartupMessage:   true,                                 // Keep logs clean
		DisableKeepalive:        false,                                // Keep connections alive (better for API performance)
		ProxyHeader:             "X-Forwarded-For",                    // Get real client IP behind proxies
		EnableTrustedProxyCheck: true,                                 // Only trust certain proxies
		TrustedProxies:          []string{"127.0.0.1", "192.168.1.1"}, // Define trusted proxies
	})
	app.Use(logger.New(logger.Config{Output: f}))
	app.Use("/api/secure/*", routes.TokenValidationAdmin)

	//app.Use("/api/users/*", routes.TokenValidationUser)

	routes.HandleRoutes(app)
	go func() {
		if err := app.Listen(":3000"); err != nil {
			log.Fatalf("Server Failed: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting Down Server ...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server ShutDown Failed: %v", err)
	}
	log.Println("Server Stopped.")

}
