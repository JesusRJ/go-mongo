package db

import (
	"reflect"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO: Cache reflection fields to tunning performance
func setField(target any, fieldName string, value any) error {
	v := reflect.ValueOf(target).Elem()
	if !v.CanAddr() {
		return ErrNotAddressable
	}
	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return ErrFieldNotFound
	}

	if _, ok := value.(primitive.ObjectID); ok {
		if field.Kind() == reflect.String {
			field.SetString(value.(primitive.ObjectID).Hex())
			return nil
		}
	}

	if value == nil {
		field.Set(reflect.Zero(v.Type()))
		return nil
	}

	field.Set(reflect.ValueOf(value))
	return nil
}

// filterWithID returns a MongoDB filter that targets a specific document by its ID,
// using the BSON format (bson.M{{"_id", ...}})
func filterWithID[T core.AbstractEntity](entity T) (bson.M, error) {
	v, ok := any(entity).(core.AbstractEntity)
	if !ok {
		return nil, ErrInvalidType
	}

	if v, ok := v.GetID().(string); ok {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return nil, err
		}
		return bson.M{"_id": id}, nil
	}
	return bson.M{"_id": v.GetID().(primitive.ObjectID)}, nil
}

func getObjectID[T core.AbstractEntity](entity T) (primitive.ObjectID, error) {
	v, ok := any(entity).(core.AbstractEntity)
	if !ok {
		return primitive.ObjectID{}, ErrInvalidType
	}

	if id, ok := v.GetID().(primitive.ObjectID); ok {
		return id, nil
	}

	return primitive.ObjectIDFromHex(v.GetID().(string))
}
