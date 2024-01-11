package redis

import (
	"context"
	"encoding/json"

	"github.com/LoomingLunar/LunarLoom-websocket-service/connection"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
)

// Read messages from channel and send to clients
func ReadMessages(ch channel) {
	// Subscribing channel
	var pubsub = connection.Redis.Subscribe(context.TODO(), string(ch))
	defer pubsub.Close()

	// Receiving messages
	var c = pubsub.Channel()
	for m := range c {
		var msg message.ClientMsg
		json.Unmarshal([]byte(m.Payload), &msg)
		if conn, exist := connections[msg.ConnectionId]; exist {
			conn <- msg.Msg
		}
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
