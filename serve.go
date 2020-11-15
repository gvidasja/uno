package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

func serveHTTP() http.HandlerFunc {
	idMatcher, err := regexp.Compile("^/(\\d+)/?$")

	if err != nil {
		log.Fatalln("Incorrect Url pattern", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		redirectToRoom := func(id string) {
			http.Redirect(w, r, fmt.Sprintf("/%s", id), http.StatusSeeOther)
		}

		serveFile := func(path string) {
			http.ServeFile(w, r, fmt.Sprintf("./client/%s", path))
		}

		if r.Method != http.MethodGet {
			http.Error(w, fmt.Sprintf("%s does not support %s method", r.URL.Path, r.Method), http.StatusMethodNotAllowed)
		}

		switch {
		case r.URL.Path == "/":
			serveFile("index.html")
		case r.URL.Path == "/new":
			redirectToRoom(strconv.Itoa(rand.Int() % 10000))
		case r.URL.Path == "/join":
			redirectToRoom(r.URL.Query().Get("roomId"))
		case idMatcher.MatchString(r.URL.Path):
			serveFile("game.html")
		default:
			serveFile(strings.TrimLeft(r.URL.Path, "/"))
		}
	}
}

func serveWs(upgrader websocket.Upgrader, rooms map[string]*Room) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)

		roomID := strings.TrimPrefix(r.URL.Path, "/ws/")

		if len(roomID) == 0 {
			http.Error(w, "Room id must be specified (e.g. /ws/69).", http.StatusBadRequest)
			return
		}

		if rooms[roomID] == nil || rooms[roomID].IsEmpty() {
			room := newRoom(roomID)
			go room.run()
			rooms[roomID] = room
		}

		room := rooms[roomID]

		if err != nil {
			log.Println("Could not upgrade", r, err)
			return
		}

		c := newConnection(ws, room)
		room.connect <- c

		go c.writeUpdates()
		go c.readActions()
	}
}
