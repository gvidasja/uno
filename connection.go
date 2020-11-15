package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	room *Room
	send chan *UnoUpdate
}

func newConnection(ws *websocket.Conn, room *Room) *Connection {
	return &Connection{
		ws:   ws,
		room: room,
		send: make(chan *UnoUpdate),
	}
}

func (c *Connection) writeUpdates() {
	defer c.ws.Close()

	for {
		update, ok := <-c.send

		if !ok {
			log.Println("Connection closed:", *c)
			c.ws.WriteMessage(websocket.CloseMessage, []byte{})
			return
		}

		c.ws.WriteJSON(update)
	}
}

func (c *Connection) readActions() {
	defer func() {
		c.room.disconnect <- c
		c.ws.Close()
	}()

	for {
		action := &UnoAction{}

		err := c.ws.ReadJSON(action)

		if err != nil {
			log.Println("Could not read message", err)
			return
		}

		log.Println("Message received:", action)

		c.room.receive <- action
	}
}
