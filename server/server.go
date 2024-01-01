package server

import (
	"os"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/internal/handlers"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

// Starting server
func Run(port string) {
	var app = fiber.New()

	// public key for jwt
	var authSecret = os.Getenv("AUTH_SECRET")

	// Check if user is logged in using jwt
	app.Use("/", jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    []byte(authSecret)},
		TokenLookup: "query:Authentication",
	}))

	// Checking if websocket upgrade possible
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Handling websocket connections
	app.Get("/ws", websocket.New(handlers.WebSocketHandler))

	app.Listen(":" + port)
}
