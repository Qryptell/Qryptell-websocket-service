package ws

import (
	"time"

	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/redis"
	"github.com/google/uuid"
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

		// Adding from,time details to the message if message is ack or user msg
		if msg.Type == message.USER_MSG || msg.Type == message.ACK_MSG {
			msg.From = c.Username
			msg.Time = time.Now().Format(time.RFC3339)

			// Sending ack message if message is user msg
			if msg.Type == message.USER_MSG {
				var id = uuid.NewString()
				var m = message.Msg{
					Id:      msg.Id,
					From:    msg.From,
					Time:    msg.Time,
					Type:    message.ACK_MSG,
					Content: message.MESSAGE_RECEIVED,
					Message: id,
				}

				c.MessageChan <- m
				msg.Id = id
			}
		}

		// Publishing messgae to redis
		go SendMsg(message.ServerMsg{ConnectionId: c.ConnectionId, Msg: msg})
	}
}

// Sending message to client
func (c *Client) WriteMsg() {
	// Getting msg from channel and sending to client response
	redis.Subscribe(c.Username, c.ConnectionId, c.MessageChan)
	for {
		select {
		case msg := <-c.MessageChan:
			var err = c.Conn.WriteJSON(msg)
			if err != nil {
				return
			}
		case _ = <-c.UnRegister:
			redis.UnSubscribe(c.ConnectionId, c.Username)
			c.Conn.Close()
			return
		}
	}
}

// Sending message to redis channel
func SendMsg(msg message.ServerMsg) {
	// Chosing message channel accoding to message type
	switch msg.Msg.Type {
	case message.USER_MSG:
		redis.PublishMsg(msg, redis.USER_MSG)
		break

	case message.ACK_MSG:
		redis.PublishMsg(msg, redis.SEND_ACK_CHANNEL)
		break

	case message.SYSTEM_MSG:
		redis.PublishMsg(msg, redis.SEND_SYSTEM_CHANNEL)
		break

	}
}
