package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadPgUrl() string {
	dir, _ := os.Getwd()
	log.Printf("Dir: %s", dir)
	err := godotenv.Load(dir + "/.env")
	if err != nil {
		log.Fatal(err)
	}

	return os.Getenv("POSTGRES_URL")
}
