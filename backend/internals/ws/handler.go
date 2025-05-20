package ws

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func SetupWebSocketRoute(app fiber.Router, hub *Hub) {
	app.Get("/votes", websocket.New(func(c *websocket.Conn) {
		hub.AddClient(c)
		defer hub.RemoveClient(c)

		for {
			if _, _, err := c.ReadMessage(); err != nil {
				break
			}
		}
	}))
}
