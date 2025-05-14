package quotes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"quotes/database"
)

type Quote struct {
	gorm.Model
	Text   string `json:"text"`
	Author string `json:"author"`
}

func GetQuote(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var quote Quote
	db.Find(&quote, id)
	return c.JSON(quote)
}

func GetQuotes(c *fiber.Ctx) error {
	db := database.DBConn
	var quotes []Quote
	db.Find(&quotes)
	return c.JSON(quotes)
}
