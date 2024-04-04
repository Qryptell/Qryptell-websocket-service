package redis

import (
	"context"

	"github.com/Qryptell/Qryptell-websocket-service/pkg/message"
	"github.com/Qryptell/Qryptell-websocket-service/util"
)

// Publishing user message to channel
func PublishMsg(msg message.ServerMsg, channel channel) error {
	// Encoding message and getting as byte string
	var data, err = util.ToByteString(msg)
	if err != nil {
		return err
	}
	// Publishing message
	err = Redis.Publish(context.TODO(), string(channel), data).Err()
	return err
}
