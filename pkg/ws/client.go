package ws

import (
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

// List of all the Clients
var Clients map[string]*Client = map[string]*Client{}

// Client structure
type Client struct {
	Conn         *websocket.Conn
	ConnectionId string
	Username     string
	MessageChan  chan (message.Msg)
	UnRegister   chan (bool)
}

// Creating new client
func NewClient(userName string, conn *websocket.Conn) *Client {
	var client = &Client{
		Conn:         conn,
		ConnectionId: uuid.NewString(),
		Username:     userName,
		MessageChan:  make(chan message.Msg),
		UnRegister:   make(chan bool),
	}

	Clients[client.ConnectionId] = client

	return client
}

// Closing connection and removing user from client list
func (c *Client) RemoveClient() {
	c.Conn.Close()
	delete(Clients, c.ConnectionId)
}
