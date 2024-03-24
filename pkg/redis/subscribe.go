package redis

import (
	"context"
	"encoding/json"

	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
)

// Read messages from channel and send to clients
func ReadMessages(ch channel) {
	// Subscribing channel
	var pubsub = Redis.Subscribe(context.TODO(), string(ch))
	defer pubsub.Close()

	// Receiving messages
	var c = pubsub.Channel()
	for m := range c {
		var msg message.ServerMsg
		json.Unmarshal([]byte(m.Payload), &msg)
		go sendMessage(msg)
	}
}

// Subscribing connection for messages
func Subscribe(userId string, connectionId string, ch chan message.Msg) {
	if connectionList, exists := connections[userId]; exists {
		connections[userId] = append(connectionList, connection{ConnectionId: connectionId, Chan: ch})
	} else {
		var c = connection{ConnectionId: connectionId, Chan: ch}
		var connectionList = []connection{c}
		connections[userId] = connectionList
	}
}

// Removing connection from connections map
func UnSubscribe(userId string, connectionId string) {
	var connectionList = connections[userId]
	for i, v := range connectionList {
		if v.ConnectionId == connectionId {
			connectionList[i] = connectionList[len(connectionList)-1]
			connectionList = connectionList[:len(connectionList)-1]
			connections[userId] = connectionList
		}
	}

	if len(connectionList) == 0 {
		delete(connections, userId)
	}
}

// Send message to all connections
func sendMessage(msg message.ServerMsg) {
	// sending messages to receivers
	for _, conn := range connections[msg.Msg.To] {
		conn.Chan <- msg.Msg
	}

	// sending message to other connection of sender
	for _, conn := range connections[msg.Msg.From] {
		if conn.ConnectionId != msg.ConnectionId {
			conn.Chan <- msg.Msg
		}
	}
}
