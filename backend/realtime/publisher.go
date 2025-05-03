package realtime

import (
	"context"
	"log"

	"github.com/PS-Wizard/Electone/internals/db"
	"github.com/gofiber/contrib/websocket"
)

func HandleLivePushUpdates(c *websocket.Conn) {
	defer c.Close()

	pubsub := db.RDB.Subscribe(context.Background(), "vote_updates")
	defer pubsub.Close()

	ch := pubsub.Channel()

	for msg := range ch {
		if err := c.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			log.Println("failed to send WS message:", err)
			break
		}
	}
}
