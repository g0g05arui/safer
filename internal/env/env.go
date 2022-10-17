package env

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	var err interface{}
	Cfg, err = godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

var Cfg map[string]string
