package db

import (
	"time"

	"github.com/jesusrj/go-mongo/core"
)

// Entity is a struct to be used as base for entities
type Entity struct {
	ID        any       `json:"id" bson:"_id,omitempty"`
	CreatedAt time.Time `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at" bson:"deleted_at,omitempty"`
}

// Ensure of Entity implements AbstractEntity
var _ core.AbstractEntity = (*Entity)(nil)

func (e Entity) GetID() any { return e.ID }
