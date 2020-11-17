package main

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/gvidasja/uno/uno"
)

type Connection struct {
	ws     *websocket.Conn
	room   *Room
	send   chan *uno.UpdateForPlayer
	player string
}

func (c *Connection) setPlayer(name string) {
	c.player = name
}

func newConnection(ws *websocket.Conn, room *Room) *Connection {
	return &Connection{
		ws:   ws,
		room: room,
		send: make(chan *uno.UpdateForPlayer),
	}
}

type Action struct {
	Action     string                 `json:"action"`
	Auth       string                 `json:"auth"`
	Data       map[string]interface{} `json:"data"`
	connection *Connection
}

func (action *Action) ToUnoAction() *uno.Action {
	return &uno.Action{
		Action: action.Action,
		Data:   action.Data,
		Player: action.Auth,
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
		action := &Action{}
		err := c.ws.ReadJSON(action)

		if err != nil {
			log.Println("Could not read message", err)
			return
		}

		log.Println("Message received:", action)

		action.connection = c

		c.room.receive <- action
	}
}
