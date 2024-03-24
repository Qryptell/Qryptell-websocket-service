package ws

import (
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
	"github.com/gofiber/contrib/websocket"
	"github.com/google/uuid"
)

// Client structure
type Client struct {
	Conn         *websocket.Conn
	ConnectionId string
	UserId       string
	MessageChan  chan (message.Msg)
	UnRegister   chan (bool)
}

// Creating new client
func NewClient(userId string, conn *websocket.Conn) *Client {
	var client = &Client{
		Conn:         conn,
		ConnectionId: uuid.NewString(),
		UserId:       userId,
		MessageChan:  make(chan message.Msg),
		UnRegister:   make(chan bool),
	}

	return client
}
