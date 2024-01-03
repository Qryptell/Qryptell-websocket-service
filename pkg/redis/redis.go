package redis

import (
	"os"

	"github.com/redis/go-redis/v9"
)

// Redis client
var Redis *redis.Client

// Type channel
type channel string

// Channel names
const USER_CHANNEL channel = "USER_MESSAGE"
const ACK_CHANNEL channel = "ACK_MESSAGE"
const SYSTEM_CHANNEL channel = "SYSTEM_MESSAGE"

// Connecting and start redis pub-sub
func Start() {
	// Getting Redis ports
	var port = os.Getenv("REDIS_IP")

	// Making connection and storing redis
	var rdb = redis.NewClient(&redis.Options{
		Addr: port,
	})
	Redis = rdb
}
