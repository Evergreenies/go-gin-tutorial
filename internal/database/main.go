package internal

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dbString := "host=localhost user=postgres password=postgres dbname=gintut port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Println("Error connection database.\n", err)

		return nil
	}

	return db
}
