package main

import "github.com/jesusrj/go-mongo/plugin/db"

type TransactionType string

type Customer struct {
	db.Entity `bson:"inline"` // Drawback: 'inline' annotation is required to prevent adding an 'entity' field instead of inner fields.
	Name      string          `bson:"name"`
	Age       uint            `bson:"age"`
	Address   string          `bson:"address"`
}

type Account struct {
	db.Entity `bson:"inline"`
	Customer  Customer `bson:"customer"`
	Bank      Bank     `bson:"bank"`
	Number    string   `bson:"number"`
	Amount    uint     `bson:"amount"`
}

type Bank struct {
	db.Entity `bson:"inline"`
	Name      string `bson:"name"`
}

type Transaction struct {
	db.Entity `bson:"inline"`
	Account   Account         `bons:"account"`
	Type      TransactionType `bons:"type"`
	Value     uint            `bons:"value"`
}

// LiteralEntity don't inherity from db.Entity
type LiteralEntity struct {
	ID    string `bson:"_id,omitempty"`
	Name  string
	Value uint
}

func (l LiteralEntity) GetID() any { return l.ID }
