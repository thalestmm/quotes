package main

import (
	"errors"
	"quotes/database"
	q "quotes/quotes"
)

type QuoteInput struct {
	Text   string
	Author string
}

func GetAllQuotes() ([]q.Quote, error) {
	db := database.DBConn
	var quotes []q.Quote
	result := db.Find(&quotes)
	return quotes, result.Error
}

func AddNewQuote(input QuoteInput) (q.Quote, error) {
	if input.Text == "" {
		return *new(q.Quote), errors.New("text cannot be blank")
	}
	if input.Author == "" {
		input.Author = "Unknown"
	}

	db := database.DBConn
	quote := q.Quote{
		Author: input.Author,
		Text:   input.Text,
	}
	result := db.Create(&quote)
	return quote, result.Error
}

func UpdateQuote(id uint, input QuoteInput) (q.Quote, error) {
	db := database.DBConn
	var quote q.Quote

	result := db.First(&quote, id)
	if result.Error != nil {
		return *new(q.Quote), result.Error
	}

	if input.Text != "" {
		quote.Text = input.Text
	}
	if input.Author != "" {
		quote.Author = input.Author
	}

	db.Save(&quote)

	return quote, nil
}

func DeleteQuote(id uint) error {
	db := database.DBConn
	var quote q.Quote
	result := db.First(&quote, id)
	if result.Error != nil {
		return result.Error
	}
	db.Delete(&quote)
	return nil
}
