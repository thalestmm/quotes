package main

import (
	"github.com/gofiber/fiber/v2"
	"quotes/quotes"
)

func SetupRoutes(app *fiber.App) {
	// FRONTEND
	app.Get("/", IndexRequest)

	// API
	api := app.Group("/api")

	// V1
	v1 := api.Group("/v1")
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"title": "V1 index",
		})
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
