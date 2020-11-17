package main

import (
	"log"

	"github.com/gvidasja/uno/uno"
)

type Room struct {
	ID          string
	connections map[*Connection]bool
	receive     chan *Action
	connect     chan *Connection
	disconnect  chan *Connection
}

func newRoom(id string) *Room {
	return &Room{
		ID:          id,
		connections: make(map[*Connection]bool),
		connect:     make(chan *Connection),
		disconnect:  make(chan *Connection),
		receive:     make(chan *Action, 256),
	}
}

func (room *Room) IsEmpty() bool {
	return len(room.connections) == 0
}

func (room *Room) run() {
	uno := uno.NewUno()

	for {
		select {
		case c := <-room.connect:
			room.connections[c] = true
			log.Println("Connected", *c)
		case c := <-room.disconnect:
			delete(room.connections, c)
			close(c.send)
			log.Println("Diconnected", *c)

			if room.IsEmpty() {
				log.Println("Room", room.ID, "empty, stopping")
				break
			}

		case action := <-room.receive:
			action.connection.player = action.Auth
			update := uno.Update(action.ToUnoAction())

			for c := range room.connections {
				c.send <- update.ForPlayer(c.player)
			}
		}
	}
}
