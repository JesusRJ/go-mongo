package tests

import "github.com/jesusrj/go-mongo/plugin/db"

const (
	// Collection names
	CollUser = "users"
	CollPets = "pets"
	CollAny  = "any"
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

// LiteralEntity don't inherity from db.Entity but implements core.AbstractEntity
type LiteralEntity struct {
	ID    string `bson:"_id,omitempty"`
	Name  string
	Value uint
}

func (l LiteralEntity) GetID() any { return l.ID }

type LiteralEntityWithoutID struct {
	Name  string
	Value uint
}

func (l LiteralEntityWithoutID) GetID() any { return "" }
