package initializers

import "github.com/MartoneWorkshop/jwt-go-test/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}