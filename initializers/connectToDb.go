package initializers

import (
	"os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectToDB() {
	var err error
	dsn := os.Getenv("DB")
	DB, err = gorm.Open(postgres.Opejn(dsn), &gorm.Config{})

	if err != nill {
		panic("Failed to connect to DB")
	}
}

