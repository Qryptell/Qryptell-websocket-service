package client

import (
	"fmt"
	"time"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/gofiber/contrib/websocket"
)

// List of all the Clients
var Clients map[string]*Client = map[string]*Client{}

// Client structure
type Client struct {
	Conn        *websocket.Conn
	SessionId   string
	Username    string
	MessageChan chan (message.UserMsg)
	AckChan     chan (message.AckMessage)
	UnRegister  chan (bool)
}

// Creating new client
func NewClient(sessionId string, userName string, conn *websocket.Conn) *Client {
	var client = &Client{
		Conn:        conn,
		SessionId:   sessionId,
		Username:    userName,
		MessageChan: make(chan message.UserMsg),
		AckChan:     make(chan message.AckMessage),
		UnRegister:  make(chan bool),
	}

	Clients[sessionId] = client

	return client
}

// Listen to messages form client
func (c *Client) ListenMsg() {
	// Disconnecting client
	defer func() {
		c.UnRegister <- true
	}()

	// Continuously listening to user messages
	for {
		// Reading client messages
		var msg message.Msg
		var err = c.Conn.ReadJSON(&msg)
		if err != nil {
			return
		}

		// Adding from,time details to the message
		msg.Message["from"] = c.Username
		msg.Message["time"] = time.Now().Format(time.RFC3339)

		// Handling message accoding to message type
		if msg.Type == message.USER_MSG {
			fmt.Println(msg)
		} else if msg.Type == message.ACK_MSG {
			fmt.Println(msg)
		} else if msg.Type == message.SYSTEM_MSG {
			fmt.Println(msg)
		}
	}
}

// Sending message to client
func (c *Client) WriteMsg() {
	// Getting msg from channel and making response
	for {
		select {
		case msg := <-c.MessageChan:
			var err = c.Conn.WriteJSON(msg)
			if err != nil {
				return
			}
		case ack := <-c.AckChan:
			var err = c.Conn.WriteJSON(ack)
			if err != nil {
				return
			}
		case _ = <-c.UnRegister:
			return
		}
	}
}

// Closing connection and removing user from client list
func (c *Client) RemoveClient() {
	c.Conn.Close()
	delete(Clients, c.SessionId)
}
