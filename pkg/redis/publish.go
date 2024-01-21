package redis

import (
	"context"

	con "github.com/LoomingLunar/LunarLoom-websocket-service/connection"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
	"github.com/LoomingLunar/LunarLoom-websocket-service/util"
)

// Publishing user message to channel
func PublishMsg(msg message.ServerMsg, channel channel) error {
	// Encoding message and getting as byte string
	var data, err = util.ToByteString(msg)
	if err != nil {
		return err
	}
	// Publishing message
	err = con.Redis.Publish(context.TODO(), string(channel), data).Err()
	return err
}
