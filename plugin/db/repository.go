package db

import (
	"context"
	"time"

	"github.com/jesusrj/go-mongo/core"
	"github.com/jesusrj/go-mongo/plugin/codec"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository[T core.Entity] struct {
	coll *mongo.Collection
	enc  *codec.Encoder
}

func NewRepository[T core.Entity](coll *mongo.Collection) (core.Repository[T], error) {
	encoder, err := codec.NewEncoder()
	if err != nil {
		return nil, err
	}

	return &AbstractRepository[T]{
		coll: coll,
		enc:  encoder,
	}, nil
}

func (a *AbstractRepository[T]) FindOne(ctx context.Context, entity *T) (*T, error) {
	filter := filterWithFields(entity)

	var result T
	if err := a.coll.FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (a *AbstractRepository[T]) FindByID(ctx context.Context, id any) (*T, error) {
	filter, err := filterWithID(Entity{ID: id})
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
	now := time.Now()
	if err := setOptionalFields(entity, map[string]any{"CreatedAt": now, "UpdatedAt": now}); err != nil {
		return nil, err
	}

	e, err := a.enc.Encode(entity)
	if err != nil {
		return nil, err
	}

	res, err := a.coll.InsertOne(ctx, e)
	if err != nil {
		return nil, err
	}

	if err := setOptionalField(entity, "ID", res.InsertedID); err != nil {
		return nil, err
	}

	return entity, nil
}

func (a *AbstractRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	if err := setOptionalFields(entity, map[string]any{"UpdatedAt": time.Now()}); err != nil {
		return nil, err
	}

	id, err := getObjectID(entity)
	if err != nil {
		return nil, err
	}

	// Preserve entity parameter
	cp := *entity

	// Set ID as primitive.ObjectID
	if err := setOptionalField(&cp, "ID", id); err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": cp}
	_, err = a.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

// TODO: Implement logical exclusion option.
func (a *AbstractRepository[T]) Delete(ctx context.Context, entity *T) (*T, error) {
	filter, err := filterWithID(*entity)
	if err != nil {
		return nil, err
	}
	_, err = a.coll.DeleteOne(ctx, filter)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// func (a *AbstractRepository[T]) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
// 	return nil
// }
