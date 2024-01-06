package redis

import (
	"context"
	"encoding/json"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/database"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
)

// Read messages from channel and send to clients
func ReadMessages(ch channel) {
	// Subscribing channel
	var pubsub = database.Redis.Subscribe(context.TODO(), string(ch))
	defer pubsub.Close()

	// Receiving messages
	var c = pubsub.Channel()
	for msg := range c {
		var m message.ClientMsg
		json.Unmarshal([]byte(msg.Payload), &m)
		go SendMsg(m)
	}
}

// Sending message to correct client
func SendMsg(msg message.ClientMsg) {
	if conn, exist := connections[msg.ConnectionId]; exist {
		conn <- msg.Msg
	}
}

// Subscribing connection for messages
func Subscribe(connectionId string, ch chan message.Msg) {
	connections[connectionId] = ch
}

// Removing connection from connections map
func UnSubscribe(connectionId string) {
	delete(connections, connectionId)
}
