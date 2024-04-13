package core

import "context"

type AbstractEntity interface {
	GetID() string
}

type AbstractRepository[T AbstractEntity] interface {
	Find(ctx context.Context, entity *T) (*T, error)
	Save(ctx context.Context, entity *T) (*T, error)
	Update(ctx context.Context, entity *T) (*T, error)
	Delete(ctx context.Context, entity *T) (*T, error)
	Tx(ctx context.Context, fn func(ctx context.Context) error) error
}

type AbstractPaginatedRepository[T any, Q any] interface {
	FindAll(ctx context.Context, query Q) (*PaginationQuery[T], error)
}
