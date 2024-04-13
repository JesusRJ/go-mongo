package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository[T any] struct {
	coll *mongo.Collection
}

func NewRepository[T any](coll *mongo.Collection) *AbstractRepository[T] {
	return &AbstractRepository[T]{
		coll: coll,
	}
}

func (a *AbstractRepository[T]) Find(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

func (a *AbstractRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

func (a *AbstractRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

func (a *AbstractRepository[T]) Delete(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

func (a *AbstractRepository[T]) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return nil
}
