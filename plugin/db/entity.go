package db

import (
	"time"

	"github.com/jesusrj/go-mongo/core"
)

// Entity is a struct to be used as base for entities
type Entity struct {
	ID        any       `bson:"_id,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty"`
	// DeletedAt time.Time `bson:"deleted_at,omitempty"` // Reserved to soft delete
}

// Ensure of Entity implements AbstractEntity
var _ core.AbstractEntity = (*Entity)(nil)

func (e Entity) GetID() any { return e.ID }
