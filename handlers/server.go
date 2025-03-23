package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/websocket"
)

// ICECandidate represents an ICE candidate from WebRTC
type ICECandidate struct {
	Candidate     string `json:"candidate"`
	SDPMid        string `json:"sdpMid"`
	SDPMLineIndex int    `json:"sdpMLineIndex"`
}

// SignalMessage represents a signaling message for WebRTC
type SignalMessage struct {
	Type string      `json:"type"`
	Data any `json:"data"`
	From string      `json:"from,omitempty"`
	To   string      `json:"to,omitempty"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true 
	},
}

// serves the main HTML page
func Homehandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("static", "index.html"))
}

// handles incoming ICE candidates from peers
func HandleICECandidate(w http.ResponseWriter, r *http.Request) {
	var candidate ICECandidate
	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		http.Error(w, "Invalid ICE candidate", http.StatusBadRequest)
		return
	}

	
	log.Printf("Received ICE candidate: %+v", candidate)

	w.WriteHeader(http.StatusOK)
}

//  handles WebSocket connections for signaling
func SignalingHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()
	

	for {
		var msg SignalMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		// Handle different types of signaling messages
		switch msg.Type {
		case "offer":
			log.Printf("Received offer: %+v", msg)
		case "answer":
			log.Printf("Received answer: %+v", msg)
		case "ice-candidate":
			log.Printf("Received ICE candidate via WebSocket: %+v", msg)
		default:
			log.Println("Unknown message type:", msg.Type)
		}
	}
}
