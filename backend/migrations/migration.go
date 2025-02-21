package main

import (
	"github.com/R-koma/translation-app/backend/infra"
	"github.com/R-koma/translation-app/backend/models"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic("Failed to migrate User database")
	}
	if err := db.AutoMigrate(&models.FriendRequest{}); err != nil {
		panic("Failed to migrate FriendRequest database")
	}
}
