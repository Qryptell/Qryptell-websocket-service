/*
				Copyright © 2023 LunarLoom 
LunarLoom Web Socket Service - WebSocket Service for LunarLoom End To End Encrypted Chat App

*/

package main

import (
	"os"

	"github.com/LoomingLunar/LunarLoom-websocket-service/connection"
	"github.com/LoomingLunar/LunarLoom-websocket-service/pkg/redis"
	"github.com/LoomingLunar/LunarLoom-websocket-service/server"
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
	connection.RedisSetUp()
	redis.ListenChannels()

	server.Run(port)
}
