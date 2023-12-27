package server

import "github.com/gofiber/fiber/v2"

// Starting server
func Run(port string) {
	var app = fiber.New()

	app.Listen(":" + port)
}
