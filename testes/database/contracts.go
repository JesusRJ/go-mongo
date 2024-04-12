package database

import "context"

// Entity represents a single instance of a domain model saved into
// the underlying database or service.
type Entity interface {
	GetID() any
}

AbstractRepositoryEntity

// Repository represents a specialized Service interface
// that provides strong-typed data access (for example, CRUD)
// operations of a domain model against the underlying database
// or service.
type Repository interface {
	Find(ctx context.Context, e *Entity) (*Entity, error)
	Save(ctx context.Context, data *Entity) error
	Update(ctx context.Context, data *Entity) error
	Delete(ctx context.Context, e *Entity) error
}

type Datasource interface {
}

type Transaction interface {
}

// set e superset 