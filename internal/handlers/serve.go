package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/gvidasja/uno/internal/room"
)

func ServeHTTP() http.HandlerFunc {
	idMatcher, err := regexp.Compile("^/(\\d+)/?$")

	if err != nil {
		log.Fatalln("Incorrect Url pattern", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		setSessionCookie(w, r)

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

func ServeWs(upgrader websocket.Upgrader, rooms map[string]*room.Room) http.HandlerFunc {
	roomIDFinder := regexp.MustCompile("/ws/?(.*)")

	return func(w http.ResponseWriter, r *http.Request) {
		cookie := setSessionCookie(w, r)

		ws, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Println("Could not upgrade", r, err)
			return
		}

		roomID := roomIDFinder.ReplaceAllString(r.URL.Path, "$1")

		if len(roomID) == 0 {
			http.Error(w, "Room id must be specified (e.g. /ws/69).", http.StatusBadRequest)
			return
		}

		if room, ok := rooms[roomID]; !ok || room.IsEmpty() {
			room := room.New(roomID)
			go room.run()
			rooms[roomID] = room
		}

		room := rooms[roomID]

		c := newConnection(ws, room, cookie)
		room.connect <- c

		go c.writeUpdates()
		go c.readActions()
	}
}

func setSessionCookie(w http.ResponseWriter, r *http.Request) string {
	var cookieValue string

	if cookie, err := r.Cookie("session"); err != nil || cookie == nil {
		cookieValue = uuid.NewString()

		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: cookieValue,
		})
	} else {
		cookieValue = cookie.Value
	}

	return cookieValue
}
