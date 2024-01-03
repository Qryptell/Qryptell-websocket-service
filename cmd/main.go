package main

import (
	"os"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/redis"
	"github.com/LoomingLunar/LunarLoom-WebSocketService/server"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")

	// Server port
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "9000"
	}

	// Connecting and creating redis client
	redis.Start()

	server.Run(port)
}
