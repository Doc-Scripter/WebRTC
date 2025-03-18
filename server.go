package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// clients holds all active WebSocket connections.
var clients = make(map[*websocket.Conn]bool)

// wsHandler handles incoming WebSocket connections and relays messages to other clients.
func wsHandler(ws *websocket.Conn) {
	clients[ws] = true
	log.Println("Client connected")
	defer func() {
		ws.Close()
		delete(clients, ws)
		log.Println("Client disconnected")
	}()

	var msg string
	for {
		// Receive a message from the connected client.
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			log.Println("Error receiving message:", err)
			break
		}
		log.Println("Received:", msg)
		// Relay the message to all other connected clients.
		for client := range clients {
			if client != ws {
				if err := websocket.Message.Send(client, msg); err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}

func main() {
	http.Handle("/ws", websocket.Handler(wsHandler))
	log.Println("Signaling server is running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
