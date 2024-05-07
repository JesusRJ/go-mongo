package tests

import "github.com/jesusrj/go-mongo/plugin/db"

const (
	// Collection names
	CollCompany = "company"
	CollUser    = "user"
	CollPet     = "pet"
	CollAny     = "any"
)

type Company struct {
	db.Entity `bson:"inline"`
	Name      string `bson:"name"`
}

type User struct {
	db.Entity `bson:"inline"`
	Name      string   `bson:"name"`
	Address   *Address `bson:"address"` // embedded
	Phone     []*Phone `bson:"phones"`  // embedded
	Company   *Company `bson:"-"`
	Pets      []*Pet   `bson:"-"`
}

type Address struct {
	Street string `bson:"street"`
	Number int    `bson:"number"`
}

type Phone struct {
	User   *User  `bson:"-"`
	Number string `bson:"number"`
}

type Pet struct {
	db.Entity `bson:"inline"`
	User      *User  `bson:"-"`
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
