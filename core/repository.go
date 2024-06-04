package core

import "context"

type FullRepository[T, Q any] interface {
	Repository[T]
	PaginatedRepository[T, Q]
}

// Entity abstract entity interface
type Entity interface {
	GetID() any
}

// Repository is an interface that defines basic CRUD (Create, Read, Update, Delete) operations
// for entities that implement the AbstractEntity interface.
type Repository[T any] interface {
	// Find retrieves an entity from the repository based on the provided context and the entity passed as a parameter.
	// It returns a pointer to the found entity and an error, if any.
	Find(ctx context.Context, entity *T) (*T, error)
	// Save saves a new entity to the repository.
	// It takes a context and a pointer to the entity to be saved.
	// It returns a pointer to the saved entity and an error, if any.
	Save(ctx context.Context, entity *T) (*T, error)
	// Update updates an existing entity in the repository.
	// It takes a context and a pointer to the entity to be updated.
	// It returns a pointer to the updated entity and an error, if any.
	Update(ctx context.Context, entity *T) (*T, error)
	// Delete removes an entity from the repository.
	// It takes a context and a pointer to the entity to be removed.
	// It returns a pointer to the removed entity and an error, if any.
	Delete(ctx context.Context, entity *T) (*T, error)

	// Tx(ctx context.Context, fn func(ctx context.Context) error) error
}

// PaginatedRepository is an interface that defines a method for retrieving paginated results
// from the repository based on a query.
type PaginatedRepository[T any, Q any] interface {
	// FindAll retrieves paginated results from the repository based on the provided context and query.
	// It takes a context and a query as parameters.
	// It returns a PaginationQuery[T] object containing the paginated results and an error, if any.
	FindAll(ctx context.Context, query Q) (*PaginationQuery[T], error)
}
