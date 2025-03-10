package main
import (
	"fmt"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)

var (
	clients      = make(map[*websocket.Conn]bool)
	broadcast    = make(chan int)
	upgrader     = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	bufferValue  = 3
	counter      = 0
	mu           sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg int
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			return
		}

		mu.Lock()
		counter += msg
		if counter%bufferValue == 0 {
			broadcast <- counter
		}
		mu.Unlock()
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
