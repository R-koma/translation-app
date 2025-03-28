package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	env := os.Getenv("ENV")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var (
		db  *gorm.DB
		err error
	)

	if env == "prod" {
		maxRetries := 10
		retryInterval := 5 * time.Second

		for i := range make([]struct{}, maxRetries) {
			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err == nil {
				log.Println("Connected to PostgresSQL database")
				break
			}
			log.Printf("Failed to connect database (attempt %d/%d): %v\n", i+1, maxRetries, err)
			time.Sleep(retryInterval)
		}
		log.Println("Setup PostgresSQL database")
	} else {
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		log.Println("Setup sqlite database")
	}

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
