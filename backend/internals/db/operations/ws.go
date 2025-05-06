package operations

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var clients = make(map[*websocket.Conn]bool)

func RegisterWebSocketRoutes(router fiber.Router) {
	router.Get("/ws/votes", websocket.New(func(c *websocket.Conn) {
		clients[c] = true
		defer func() {
			delete(clients, c)
			c.Close()
		}()

		for {
			_, _, err := c.ReadMessage() // Keep it alive
			if err != nil {
				break
			}
		}
	}))
}

func BroadcastVoteUpdate(vote Vote) {
	for client := range clients {
		err := client.WriteJSON(vote)
		if err != nil {
			client.Close()
			delete(clients, client)
		}
	}
}
