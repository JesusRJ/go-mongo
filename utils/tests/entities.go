package tests

import (
	"time"

	"github.com/jesusrj/go-mongo/plugin/db"
)

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
	Name      string    `bson:"name"`
	Address   *Address  `bson:"address"` // embedded
	Phone     []*Phone  `bson:"phones"`  // embedded
	Company   *Company  `bson:"company" ref:"belongsTo,company,company_id,_id,user"`
	Pets      []*Pet    `bson:"pets"    ref:"hasMany,pet,_id,user_id,pets"`
	Login     time.Time `bson:"login_time"`
}

type Address struct {
	Street string `bson:"street"`
	Number int    `bson:"number"`
}

type Phone struct {
	Number string `bson:"number"`
}

type Pet struct {
	db.Entity `bson:"inline"`
	User      User   `bson:"-" ref:"belongsTo,user,user_id,_id,user"`
	Name      string `bson:"name"`
}

// RegularEntity don't inherity from db.Entity but implements core.AbstractEntity
type RegularEntity struct {
	ID    string `bson:"_id,omitempty"`
	Name  string
	Value uint
}

func (l RegularEntity) GetID() any { return l.ID }

type RegularEntityWithoutID struct {
	Name  string
	Value uint
}

func (l RegularEntityWithoutID) GetID() any { return "" }
