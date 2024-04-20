package tests

import "github.com/jesusrj/go-mongo/plugin/db"

const (
	// Collection names
	CollUser = "users"
	CollPets = "pets"
)

type User struct {
	db.Entity `bson:"inline"`
	Name      string `bson:"name"`
	Address   string `bson:"address"`
	Pets      []*Pet `bson:"pets"`
}

type Pet struct {
	db.Entity `bson:"inline"`
	Name      string `bson:"name"`
}
