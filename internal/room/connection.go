package room

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/gvidasja/uno/internal/uno"
)

type connection struct {
	ws        *websocket.Conn
	room      *Room
	send      chan *uno.UpdateForPlayer
	sessionId string
}

func newConnection(ws *websocket.Conn, room *Room, sessionId string) *connection {
	return &connection{
		ws:        ws,
		room:      room,
		sessionId: sessionId,
		send:      make(chan *uno.UpdateForPlayer),
	}
}

type Action struct {
	Action     string                 `json:"action"`
	Data       map[string]interface{} `json:"data"`
	connection *connection
}

func (action *Action) ToUnoAction() *uno.Action {
	return &uno.Action{
		Action: uno.UnoAction(action.Action),
		Data:   action.Data,
		Player: action.connection.sessionId,
	}
}

func (c *connection) writeUpdates() {
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

func (c *connection) readActions() {
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
