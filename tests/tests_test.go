package tests_test

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/jesusrj/go-mongo/plugin/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	Database.Collection(CollUser).Drop(ctx)
	Database.Collection(CollCompany).Drop(ctx)
	Database.Collection(CollPet).Drop(ctx)

	repoCompany := db.NewRepository[Company](Database.Collection(CollCompany))
	repoUser := db.NewRepository[User](Database.Collection(CollUser))
	repoPet := db.NewRepository[Pet](Database.Collection(CollPet))

	cId, _ := primitive.ObjectIDFromHex(StaticCompanyID[0])
	company := &Company{Entity: db.Entity{ID: cId}, Name: "My Petshop"}
	repoCompany.Save(ctx, company)

	for x, v := range StaticUserID {
		if id, err := primitive.ObjectIDFromHex(v); err == nil {
			u, _ := repoUser.Save(ctx, GetUser("user_"+strconv.Itoa(x), Config{ID: id, Pets: 2, Company: company}))
			for _, p := range u.Pets {
				p.User = u
				repoPet.Save(ctx, p)
			}
		}
	}

	// insert data for batch tests
	// for x := range 500 {
	// 	repoUsers.Save(ctx, GetUser("user_batch_"+strconv.Itoa(x), Config{}))
	// }
}
