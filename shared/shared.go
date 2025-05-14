package shared

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"quotes/database"
	"quotes/quotes"
)

// Shared functionalities will be included here to avoid circular dependencies

func InitDatabase() {
	var err error
	db := database.DBConn

	db, err = gorm.Open(sqlite.Open("quotes.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	database.DBConn = db
	log.Println("Connection opened to database")
	db.AutoMigrate(&quotes.Quote{})
	log.Println("Database migrated")
}
