package main

import (
	"log"

	data "github.com/Vishal-2029/Quiz"
	ws "github.com/Vishal-2029/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New()

	// Serve static files (HTML)
	app.Static("/", "/frontEnd")

	// WebSocket route
		app.Get("/ws", websocket.New(ws.HandleWebSocket))

	// API to get quiz questions
	app.Get("/api/questions", func(c *fiber.Ctx) error {
		return c.JSON(data.GetQuestions())
	})

	log.Println("Server started on http://localhost:3030")
	log.Fatal(app.Listen(":3030"))
}
