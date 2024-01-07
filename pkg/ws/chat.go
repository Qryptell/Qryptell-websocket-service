package ws

import (
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/redis"
)

// Listen to messages form client
func (c *Client) ListenMsg() {
	// Disconnecting client
	defer func() {
		c.Conn.Close()
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
		// Publishing messgae to redis
		go SendMsg(msg)
	}
}

// Sending message to client
func (c *Client) WriteMsg() {
	// Getting msg from channel and sending to client response
	redis.Subscribe(c.ConnectionId, c.MessageChan)
	for {
		select {
		case msg := <-c.MessageChan:
			var err = c.Conn.WriteJSON(msg)
			if err != nil {
				return
			}
		case _ = <-c.UnRegister:
			redis.UnSubscribe(c.ConnectionId)
			c.Conn.Close()
			return
		}
	}
}

// Sending message to redis channel
func SendMsg(msg message.Msg) {
	// Chosing message channel accoding to message type
	switch msg.Type {
	case message.USER_MSG:
		redis.PublishMsg(msg, redis.SEND_USER_CHANNEL)
		break

	case message.ACK_MSG:
		redis.PublishMsg(msg, redis.SEND_ACK_CHANNEL)
		break

	case message.SYSTEM_MSG:
		redis.PublishMsg(msg, redis.SEND_SYSTEM_CHANNEL)
		break

	}
}
