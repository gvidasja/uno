package main

import "log"

type Room struct {
	ID          string
	connections map[*Connection]bool
	broadcast   chan *UnoUpdate
	receive     chan *UnoAction
	connect     chan *Connection
	disconnect  chan *Connection
}

func newRoom(id string) *Room {
	return &Room{
		ID:          id,
		connections: make(map[*Connection]bool),
		connect:     make(chan *Connection),
		disconnect:  make(chan *Connection),
		broadcast:   make(chan *UnoUpdate),
		receive:     make(chan *UnoAction, 256),
	}
}

func (room *Room) IsEmpty() bool {
	return len(room.connections) == 0
}

func (room *Room) run() {
	uno := NewUno()

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
			update := uno.Update(action)
			room.broadcast <- update
		case update := <-room.broadcast:
			for c := range room.connections {
				c.send <- update
			}
		}
	}
}
