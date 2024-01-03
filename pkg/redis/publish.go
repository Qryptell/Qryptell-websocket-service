package redis

import (
	"context"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/message"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/util"
)

// Publishing user message to channel
func PublishUserMsg(msg message.Msg) error {
	// Encoding message and getting as byte string
	var data, err = util.ToByteString(msg)
	if err != nil {
		return err
	}
	// Publishing message
	err = Redis.Publish(context.TODO(), string(USER_CHANNEL), data).Err()
	return err
}
