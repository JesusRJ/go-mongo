package db

import (
	"context"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractPaginatedRepository[T core.AbstractEntity, Q interface{}] struct {
	coll *mongo.Collection
}

func NewPaginatedRepository[T core.AbstractEntity, Q any](coll *mongo.Collection) core.AbstractPaginatedRepository[T, Q] {
	return &AbstractPaginatedRepository[T, Q]{
		coll: coll,
	}
}

func (a *AbstractPaginatedRepository[T, Q]) FindAll(ctx context.Context, query Q) (*core.PaginationQuery[T], error) {
	// c, err := a.coll.Find(ctx, query, &options.FindOptions{})
	// if err != nil {
	// 	return nil, err
	// }

	// var result []T
	// c.All(ctx, &result)

	// return &result, nil
	return nil, nil
}
