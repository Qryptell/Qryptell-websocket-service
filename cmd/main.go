package main

import (
	"os"

	"github.com/LunarLoom/WebSocketService/server"
)

func main() {
	var port string

	if port = os.Getenv("PORT"); port == ""  {
		port = "9000"
	}

	server.Run(port)
}
