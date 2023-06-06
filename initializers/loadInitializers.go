package initializers

import (
	"log"

	"github.com/joho/godotenv"
)
func LoadEnvVaraiables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}
}