package server

import (
	"os"

	"github.com/LoomingLunar/LunarLoom-WebSocketService/pkg/client"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Getting sessionId,username from jwt token
		var user = c.Locals("user").(*jwt.Token)
		var claims = user.Claims.(jwt.MapClaims)
		var username = claims["userName"].(string)
		var sessionId = claims["sessionId"].(string)

		// Creating new client and reading and writing messages
		var cli = client.NewClient(sessionId, username, c)
		go cli.ListenMsg()
		cli.WriteMsg()

		// Disconnecting connection
		cli.RemoveClient()
	}))

	app.Listen(":" + port)
}
