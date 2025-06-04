package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Solo intentamos cargar el archivo .env si existe
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}
