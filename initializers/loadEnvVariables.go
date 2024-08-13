package initializers
import (
	"log"
	"github.com/joho/godotenv"
)
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nill {
		log.Fatal("Error loading .env File")
	}
}

