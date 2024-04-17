package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jesusrj/go-mongo/plugin/db"
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

	database := dbClient.Database("XBank")

	customerRepository := db.NewRepository[Customer](database.Collection("customers"))

	c, err := customerRepository.Save(ctx, &Customer{
		Name:    "Mario Pachos",
		Age:     43,
		Address: "Sevita Street, Number 2",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", c)

	// Use a entity no inherit from core.Entity
	literalRepository := db.NewRepository[LiteralEntity](database.Collection("literal"))
	l, err := literalRepository.Save(ctx, &LiteralEntity{
		Name:  "Teste123",
		Value: 1500,
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", l)

}
