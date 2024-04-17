package db

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AbstractRepository[T any] struct {
	coll *mongo.Collection
}

func NewRepository[T core.AbstractEntity](coll *mongo.Collection) core.AbstractRepository[T] {
	return &AbstractRepository[T]{
		coll: coll,
	}
}

func (a *AbstractRepository[T]) Find(ctx context.Context, entity *T) (*T, error) {
	v, ok := any(entity).(interface{ GetID() any })
	if !ok {
		return nil, ErrInvalidType
	}

	id, err := primitive.ObjectIDFromHex(fmt.Sprintf("%v", v.GetID()))
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": id}

	var result *T
	if err := a.coll.FindOne(ctx, filter).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
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

	if err := setField(entity, "ID", res.InsertedID.(primitive.ObjectID).Hex()); err != nil {
		return nil, err
	}

	return entity, nil
}

func (a *AbstractRepository[T]) Update(ctx context.Context, entity *T) (*T, error) {
	v, ok := any(entity).(core.AbstractEntity)
	if !ok {
		return nil, ErrInvalidType
	}

	v.GetID()

	return nil, nil
}

func (a *AbstractRepository[T]) Delete(ctx context.Context, entity *T) (*T, error) {
	return nil, nil
}

// func (a *AbstractRepository[T]) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
// 	return nil
// }

// TODO: Cache reflection fields to tunning performance
func setField(target any, fieldName string, value any) error {
	v := reflect.ValueOf(target).Elem()
	if !v.CanAddr() {
		return ErrNotAddressable
	}

	field := reflect.Indirect(v).FieldByName(fieldName)
	if !field.IsValid() {
		return ErrFieldNotFound
	}

	field.Set(reflect.ValueOf(value))
	return nil
}
