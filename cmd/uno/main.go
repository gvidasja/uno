package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/gvidasja/uno/internal/handlers"
)

func main() {
	port := Getenv("PORT", "3000")

	rooms := make(map[string]*Room)
	upgrader := websocket.Upgrader{}

	http.HandleFunc("/ws/", handlers.ServeWs(upgrader, rooms))
	http.HandleFunc("/", handlers.ServeHTTP())

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatalln("Could not start server: ", err)
	}
}

func Getenv(key string, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
