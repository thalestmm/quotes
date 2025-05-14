package quotes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"quotes/database"
	"time"
)

type Quote struct {
	gorm.Model
	Text   string `json:"text"`
	Author string `json:"author"`
}

// TODO: Check if quotes have "" already or not

func GetQuote(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var quote Quote
	db.Find(&quote, id)
	return c.JSON(quote)
}

func GetRandomQuote(c *fiber.Ctx) error {
	db := database.DBConn
	id := GetRandomID()
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

func GetRandomID() uint {
	var ids []uint
	db := database.DBConn
	if err := db.Model(&Quote{}).Pluck("id", &ids).Error; err != nil {
		panic(err)
	}
	if len(ids) == 0 {
		log.Fatal("No quotes found")
	}

	r := rand.New(rand.NewSource(time.Now().UTC().UnixNano())) // TODO: Change to current day

	randomIndex := r.Intn(len(ids))
	randomID := ids[randomIndex]

	return randomID
}
