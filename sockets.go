package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Maintain a map of documentID to list of WebSocket connections
var documentConnections = make(map[string][]*websocket.Conn)

func HandleWebSocket(c *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Get documentID from URL parameters
	documentID := c.Param("id")

	// Add connection to list of connections for this document
	documentConnections[documentID] = append(documentConnections[documentID], conn)

	// Read and broadcast messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from WebSocket:", err)
			break
		}

		// Broadcast message to all connections for this document
		for _, c := range documentConnections[documentID] {
			if c != conn {
				err := c.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Println("Error writing message to WebSocket:", err)
				}
			}
		}
	}
}
