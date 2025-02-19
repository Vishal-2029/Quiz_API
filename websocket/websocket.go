package ws

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	
)

var clients = make(map[*websocket.Conn]bool)

// Handle WebSocket connection
func HandleWebSocket(c *websocket.Conn) {
	defer func() {
		delete(clients, c)
		c.Close()
	}()

	clients[c] = true
	log.Println("New WebSocket Connection")

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket Error:", err)
			break
		}

		fmt.Println("Received:", string(msg))

		// Broadcast to all clients
		for client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Println("Error writing WebSocket message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
