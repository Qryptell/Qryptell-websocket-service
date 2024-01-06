package redis

import (
	"context"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/database"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/util"
)

// Publishing user message to channel
func PublishMsg(msg message.Msg, channel channel) error {
	// Encoding message and getting as byte string
	var data, err = util.ToByteString(msg)
	if err != nil {
		return err
	}
	// Publishing message
	err = database.Redis.Publish(context.TODO(), string(channel), data).Err()
	return err
}
