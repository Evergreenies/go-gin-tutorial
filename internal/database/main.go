package internal

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbString := os.Getenv("DB_CONNECTION_STRING")
	if dbString == "" {
		log.Fatalln("Make sure you have provided database connection string in configurations")

		return nil
	}

	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Println("Error connection database.\n", err)

		return nil
	}

	return db
}
