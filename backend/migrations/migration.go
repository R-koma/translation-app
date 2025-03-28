package main

import (
	"github.com/R-koma/translation-app/backend/db"
	"github.com/R-koma/translation-app/backend/models"
)

func main() {
	db := db.SetupDB()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate User database")
	}
	if err := db.AutoMigrate(&models.FriendRequest{}); err != nil {
		panic("Failed to migrate FriendRequest database")
	}
}
