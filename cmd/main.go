/*
				Copyright © 2023 Qryptell
Qryptell Web Socket Service - WebSocket Service for Qryptell End To End Encrypted Chat App

*/

package main

import (
	"os"

	"github.com/Qryptell/Qryptell-websocket-service/pkg/redis"
	"github.com/Qryptell/Qryptell-websocket-service/server"
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
	redis.SetUp()

	server.Run(port)
}
