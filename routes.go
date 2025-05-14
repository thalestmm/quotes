package main

import (
	"github.com/gofiber/fiber/v2"
	"quotes/quotes"
)

func setupRoutes(app *fiber.App) {
	// FRONTEND
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// API
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Index")
	})

	// Quotes functionality
	q := v1.Group("/quotes")
	q.Get("/", func(c *fiber.Ctx) error {
		return quotes.GetQuotes(c)
	})
	q.Get("/random", func(c *fiber.Ctx) error {
		return quotes.GetRandomQuote(c)
	})
	q.Get("/:id", func(c *fiber.Ctx) error {
		return quotes.GetQuote(c)
	})
}
