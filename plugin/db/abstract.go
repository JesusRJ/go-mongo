package db

import (
	"context"
	"time"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository[T core.AbstractEntity] struct {
	coll *mongo.Collection
}

func NewRepository[T core.AbstractEntity](coll *mongo.Collection) core.AbstractRepository[T] {
	return &AbstractRepository[T]{
		coll: coll,
	}
}

func (a *AbstractRepository[T]) Find(ctx context.Context, entity *T) (*T, error) {
	filter, err := filterWithID(*entity)
	if err != nil {
		return nil, err
	}

	var result T
	if err := a.coll.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *AbstractRepository[T]) Save(ctx context.Context, entity *T) (*T, error) {
	if err := setField(entity, "CreatedAt", time.Now()); err != nil && err != ErrFieldNotFound {
		return nil, err
	}
	if err := setField(entity, "UpdatedAt", time.Now()); err != nil && err != ErrFieldNotFound {
		return nil, err
	}

	res, err := a.coll.InsertOne(ctx, entity)
	if err != nil {
		return nil, err
	}

	if err := setField(entity, "ID", res.InsertedID); err != nil {
		return nil, err
	}

	return entity, nil
}

func (a *AbstractRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	if err := setField(entity, "UpdatedAt", time.Now()); err != nil && err != ErrFieldNotFound {
		return nil, err
	}

	filter, err := filterWithID(*entity)
	if err != nil {
		return nil, err
	}

	// remove the ID field
	setField(entity, "ID", primitive.NilObjectID)

	update := bson.M{"$set": entity}

	_, err = a.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (a *AbstractRepository[T]) Delete(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

// func (a *AbstractRepository[T]) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
// 	return nil
// }
