package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"quotes/shared"
)

func main() {
	app := fiber.New()

	setupRoutes(app)
	shared.InitDatabase()

	log.Fatal(app.Listen(":3000"))
}
