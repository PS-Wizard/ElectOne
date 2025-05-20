package ws

import (
	"sync"

	"github.com/gofiber/websocket/v2"
)

type Hub struct {
	clients map[*websocket.Conn]bool
	lock    sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) AddClient(c *websocket.Conn) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.clients[c] = true
}

func (h *Hub) RemoveClient(c *websocket.Conn) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.clients, c)
}

func (h *Hub) Broadcast(message []byte) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for c := range h.clients {
		if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
			c.Close()
			delete(h.clients, c)
		}
	}
}
