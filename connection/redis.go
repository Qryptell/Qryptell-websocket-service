package connection

import (
	"os"

	"github.com/redis/go-redis/v9"
)

// Redis client
var Redis *redis.Client

// Connecting and start redis pub-sub
func RedisSetUp() {
	// Getting Redis ports
	var port = os.Getenv("REDIS_IP")

	// Making connection and storing redis
	var rdb = redis.NewClient(&redis.Options{
		Addr: port,
	})
	Redis = rdb
}
