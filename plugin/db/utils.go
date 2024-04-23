package db

import (
	"errors"
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

func setFields(target any, fields map[string]any) error {
	for k, v := range fields {
		if err := setField(target, k, v); err != nil && !errors.Is(err, ErrFieldNotFound) {
			return err
		}
	}
	return nil
}

// filterWithID returns a MongoDB filter that targets a specific document by its ID,
// using the BSON format (bson.M{{"_id", ...}})
func filterWithID[T core.AbstractEntity](entity T) (bson.M, error) {
	id, err := getObjectID(entity)
	if err != nil {
		return nil, err
	}
	return bson.M{"_id": id}, nil
}

// getObjectID returns the ObjectID associated with the provided entity.
// If the entity does not implement the core.AbstractEntity interface,
// it returns an ErrInvalidType error.
// If the entity implements the interface, it checks if the ID is an ObjectID.
// If it is, it returns the ObjectID. Otherwise, it attempts to convert
// the ID to an ObjectID from its hexadecimal representation.
// Returns the ObjectID and an error, if any.
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
