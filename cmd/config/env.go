package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
