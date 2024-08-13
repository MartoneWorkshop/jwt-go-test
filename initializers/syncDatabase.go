package initializers

import "project2/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}