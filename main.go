package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	port := Getenv("PORT", "3000")

	rooms := make(map[string]*Room)
	upgrader := websocket.Upgrader{}

	http.HandleFunc("/ws/", serveWs(upgrader, rooms))
	http.HandleFunc("/", serveHTTP())

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalln("Could not start server: ", err)
	}
}
