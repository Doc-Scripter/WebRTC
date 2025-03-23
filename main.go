package main

import (
	"net/http"
	"webRTC/handlers"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", handlers.Homehandler)

	mux.HandleFunc("/ice-candidate", handlers.HandleICECandidate)

	mux.HandleFunc("/ws", handlers.SignalingHandler)

	http.ListenAndServe("Running on localhost:8080", mux)
}
