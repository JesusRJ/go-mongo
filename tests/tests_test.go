package tests_test

import (
	"context"
	"log"
	"os"

	"github.com/jesusrj/go-mongo/plugin/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

const mongDSN = "mongodb://root:MongoPass321!@localhost:27017"

var Database *mongo.Database

func init() {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(mongDSN)

	Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("error connecting to database, got error: %+v", err)
		os.Exit(1)
	}

	err = Client.Ping(ctx, nil)
	if err != nil {
		log.Printf("failed pinging database, got error: %+v", err)
		os.Exit(1)
	}

	Database = Client.Database("petshop")

	seed()
}

// Populate database with tests values
func seed() {
	ctx := context.Background()

	repoUsers := db.NewRepository[User](Database.Collection(CollUser))

	for _, id := range StaticID {
		repoUsers.Save(ctx, GetUser("", Config{ID: id}))
	}
}
