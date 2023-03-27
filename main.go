package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jesusrj/go-mongo/config"
	"github.com/jesusrj/go-mongo/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.New(
		config.Global().MongoDbHost,
		config.Global().MongoDbUsername,
		config.Global().MongoDbPassword,
		config.Global().MongoDbPort)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = db.Connect(ctx)
	if err != nil {
		log.Fatalf("error connecting to database: %+v", err)
	}

	fmt.Printf("%+v\n%+v\n", config.Global(), db)
}
