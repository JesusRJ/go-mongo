package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("error connecting to database: %+v", err)
	}

	err = dbClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("error pinging database: %+v", err)
	}

	fmt.Println("Database connected")
}
