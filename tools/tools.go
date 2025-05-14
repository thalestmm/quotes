package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"quotes/database"
	q "quotes/quotes"
	"quotes/shared"
	"strings"
)

// Internal admin tooling

func main() {
	shared.InitDatabase()
	reader := bufio.NewReader(os.Stdin)

	var quoteText string
	var quoteAuthor string

	fmt.Println("Type the quote text: ")
	// Read all words from the current terminal line
	quoteText, _ = reader.ReadString('\n')
	quoteText = strings.TrimSpace(quoteText)

	fmt.Println("\nType the quote author: ")
	quoteAuthor, _ = reader.ReadString('\n')
	quoteAuthor = strings.TrimSpace(quoteAuthor)

	input := QuoteInput{
		Text:   quoteText,
		Author: quoteAuthor,
	}

	quote, err := AddNewBook(input)
	if err != nil {
		panic(err)
	}
	log.Printf("[Object created] ID: %d | Text: %s | Author: %s", quote.ID, quote.Text, quote.Author)
}

type QuoteInput struct {
	Text   string
	Author string
}

func AddNewBook(input QuoteInput) (q.Quote, error) {
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
