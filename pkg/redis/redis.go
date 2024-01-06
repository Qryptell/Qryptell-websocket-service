package redis

import (
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
)

// Type channel
type channel string

// Channel names
const SEND_USER_CHANNEL channel = "USER_MESSAGE"
const SEND_ACK_CHANNEL channel = "ACK_MESSAGE"
const SEND_SYSTEM_CHANNEL channel = "SYSTEM_MESSAGE"
const RECEIVE_CHAT_CHANNEL channel = "CHAT_MESSAGE"

// Storing all connections
var connections = make(map[string]chan<- message.Msg)

// Listening to redis channels
func ListenChannels() {
	go ReadMessages(RECEIVE_CHAT_CHANNEL)
}
