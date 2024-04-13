package core

import "time"

type Entity struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// Ensure of Entity implements AbstractEntity
var _ AbstractEntity = (*Entity)(nil)

func (e *Entity) GetID() string { return e.ID }
