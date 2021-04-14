package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

const (
	dataTypeString = "STRING"
	dataTypeTime   = "TIME"
)

func showHelp() {
	fmt.Println("USAGE:\n\tapps [arguments...]")
	fmt.Println("COMMANDS:")
	fmt.Println("\tSET [redis-key] [data-type] [value]")
	fmt.Println("\tGET [redis-key] [data-type]")
	fmt.Println("DATA TYPES:")
	fmt.Println("\tSTRING escape strings")
	fmt.Println("\tTIME RFC3339 format, ex: 2021-04-14T08:09:47Z")
	os.Exit(0)
}

func validateLenArgs(i int) {
	if len(os.Args) < i {
		showHelp()
	}
}

func validateDataType(d string) {
	supportedDataType := map[string]bool{
		dataTypeString: true,
		dataTypeTime:   true,
	}

	if !supportedDataType[d] {
		showHelp()
	}
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		showHelp()
	}

	// Validation Args
	switch os.Args[1] {
	case "SET":
		validateLenArgs(5)
	case "GET":
		validateLenArgs(4)
	default:
		showHelp()
	}

	dataType := os.Args[3]
	redisKey := os.Args[2]
	validateDataType(dataType)

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password:     os.Getenv("REDIS_PASSWORD"),
		DB:           db,
		PoolSize:     64,
		PoolTimeout:  10,
		MinIdleConns: 1,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "GET":
		switch dataType {
		case dataTypeTime:
			var t time.Time
			if err := redisClient.Get(redisKey).Scan(&t); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			fmt.Println(t.Format(time.RFC3339))
		case dataTypeString:
			var s string
			if err := redisClient.Get(redisKey).Scan(&s); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			s = strings.Replace(s, `"`, `\"`, -1)

			fmt.Printf(`"%s"`, s)
			fmt.Println()
		}

	case "SET":
		value := os.Args[4]
		switch dataType {
		case dataTypeString:
			if err := redisClient.Set(redisKey, value, 0).Err(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		case dataTypeTime:
			t, err := time.Parse(time.RFC3339, value)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			if err := redisClient.Set(redisKey, t, 0).Err(); err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
		}
	}

	redisClient.Close()
}
