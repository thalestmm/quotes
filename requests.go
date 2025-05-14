package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
)

func JSONFromEndpoint(c *fiber.Ctx, endpoint string) (map[string] interface, error) {
	// Make a request to the JSON API
	url := "http://localhost:3000/api/v1" + endpoint
	resp, err := http.Get(url)
	if err != nil {
		return nil, c.Status(500).SendString("Failed to fetch user data")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("Failed to close response body: ", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, c.Status(500).SendString("Failed to read response")
	}
	var quoteData map[string]interface{}
	if err := json.Unmarshal(body, &quoteData); err != nil {
		return nil, c.Status(500).SendString("Failed to parse JSON")
	}
	return quoteData, nil
}

func IndexRequest(c *fiber.Ctx) error {
	quoteData, err := JSONFromEndpoint(c, "/quotes/random")
	if err != nil {
		log.Println(err)
	}
	// Now render the page with the fetched data
	return c.Render("index", quoteData)
}
