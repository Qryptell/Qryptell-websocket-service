package ws

import (
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

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

	return client
}
