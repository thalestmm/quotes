package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"quotes/shared"
	"strings"
)

// Internal admin tooling

func main() {
	shared.InitDatabase()
	reader := bufio.NewReader(os.Stdin)

	var quoteText string
	var quoteAuthor string
	var quoteID uint

	fmt.Println("\nSelect mode: (N) New quote | (U) Update quote | (D) Delete quote | (L) List all quotes | (Q) Quit")
	var mode string
	_, _ = fmt.Scanln(&mode)
	mode = strings.ToLower(mode)

	// TODO: Add a context loop

	if mode == "q" {
		os.Exit(0)
	}

	if mode == "l" {
		quotes, err := GetAllQuotes()
		if err != nil {
			log.Fatal(err)
		}
		for _, quote := range quotes {
			fmt.Printf("ID: %d | Text: %s | Author: %s\n", quote.ID, quote.Text, quote.Author)
		}
	}

	if mode == "d" || mode == "u" {
		fmt.Println("Select the quote ID: ")
		_, err := fmt.Scanln(&quoteID)
		if err != nil {
			log.Fatal(err)
		}
		if mode == "d" {
			err := DeleteQuote(quoteID)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if mode == "u" || mode == "n" {
		fmt.Println("\nType the quote text: ")
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
		if mode == "n" {
			quote, err := AddNewQuote(input)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[Object created] ID: %d | Text: %s | Author: %s", quote.ID, quote.Text, quote.Author)
		}
		if mode == "u" {
			quote, err := UpdateQuote(quoteID, input)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("[Object updated] ID: %d | Text: %s | Author: %s", quote.ID, quote.Text, quote.Author)
		}
	}
}
