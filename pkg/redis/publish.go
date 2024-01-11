package redis

import (
	"context"

	"github.com/LoomingLunar/LunarLoom-websocket-service/connection"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/message"
	"github.com/LoomingLunar/LunarLoom-websocket-service/util"
)

// Publishing user message to channel
func PublishMsg(msg message.Msg, channel channel) error {
	// Encoding message and getting as byte string
	var data, err = util.ToByteString(msg)
	if err != nil {
		return err
	}
	// Publishing message
	err = connection.Redis.Publish(context.TODO(), string(channel), data).Err()
	return err
}
