package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"simple-redis-cli/src"
)

func main() {
	fmt.Println("redis-simple-cli")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	src.InitRedis()
}
