package redis

import (
	"github.com/Qryptell/Qryptell-websocket-service/pkg/message"
)

// Type channel
type channel string

// Channel names
const USER_MSG channel = "USER_MESSAGE"
const SEND_ACK_CHANNEL channel = "ACK_MESSAGE"
const SEND_SYSTEM_CHANNEL channel = "SYSTEM_MESSAGE"

// Storing all connections
var connections = make(map[string][]connection)

type connection struct{
	ConnectionId string
	Chan chan <- message.Msg
}

// Listening to redis channels
func ListenChannels() {
	go ReadMessages("USER_MESSAGE")
}
