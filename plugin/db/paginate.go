package db

import (
	"context"

	"github.com/jesusrj/go-mongo/core"
	"github.com/jesusrj/go-mongo/plugin/codec"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AbstractPaginatedRepository[T core.Entity, Q interface{}] struct {
	coll *mongo.Collection
	enc  *codec.Encoder
}

func NewPaginatedRepository[T core.Entity, Q any](coll *mongo.Collection) (core.PaginatedRepository[T, Q], error) {
	encoder, err := codec.NewEncoder()
	if err != nil {
		return nil, err
	}

	return &AbstractPaginatedRepository[T, Q]{
		coll: coll,
		enc:  encoder,
	}, nil
}

func (a *AbstractPaginatedRepository[T, Q]) FindAll(ctx context.Context, query Q, opts core.QueryOptions) (*core.Pagination[T], error) {
	count, err := a.coll.CountDocuments(ctx, query)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &core.Pagination[T]{Total: 0, Data: nil}, nil
	}

	c, err := a.coll.Find(ctx, query, findOptions(opts))
	if err != nil {
		return nil, err
	}

	var result []T
	if err := c.All(ctx, &result); err != nil {
		return nil, err
	}

	return &core.Pagination[T]{
		Page:  int(*opts.Page),
		Pages: int(count) / int(*opts.PageSize),
		Total: int(count),
		Data:  result,
	}, nil
}

func findOptions(opts core.QueryOptions) *options.FindOptions {
	o := options.Find()

	if opts.PageSize != nil {
		o.SetLimit(*opts.PageSize)
	}

	if opts.Page != nil {
		o.SetSkip((*opts.PageSize - 1) * *opts.PageSize)
	}

	if opts.Order != nil {
		o.SetSort(bson.D{{Key: opts.Order.Field, Value: opts.Order.Direction}})
	}

	return o
}
