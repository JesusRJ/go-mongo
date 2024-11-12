package db

import (
	"errors"
	"reflect"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var NilObjectID = primitive.NilObjectID

// TODO: Cache reflection fields to tunning performance

func setField(target any, fieldName string, value any) error {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return ErrInvalidTarget
	}

	v = v.Elem()
	if !v.CanAddr() {
		return ErrNotAddressable
	}

	field := v.FieldByName(fieldName)
	if !field.IsValid() {
		return ErrFieldNotFound
	}
	if !field.CanSet() {
		return ErrFieldCannotBeSet
	}

	if _, ok := value.(primitive.ObjectID); ok {
		if field.Kind() == reflect.String {
			field.SetString(value.(primitive.ObjectID).Hex())
			return nil
		}
		if field.Kind() == reflect.Interface {
			field.Set(reflect.ValueOf(value))
			return nil
		}
	}

	val := reflect.ValueOf(value)
	if field.Type() != val.Type() {
		return ErrIncorrectTypeForField
	}

	field.Set(val)
	return nil
}

// setOptionalField set the field on the target object, ignoring if field do not exist.
func setOptionalField(target any, fieldName string, value any) error {
	if err := setField(target, fieldName, value); err != nil && !errors.Is(err, ErrFieldNotFound) {
		return err
	}
	return nil
}

// setOptionalFields sets each field on the target object, ignoring fields that do not exist.
// It iterates through the fields map and calls the setField function to set each field on the target object.
// If an error occurs while setting a field that does not exist, it is ignored.
func setOptionalFields(target any, fields map[string]any) error {
	for k, v := range fields {
		if err := setOptionalField(target, k, v); err != nil {
			return err
		}
	}
	return nil
}

// getObjectID returns the ObjectID associated with the provided entity.
// If the entity does not implement the core.AbstractEntity interface,
// it returns an ErrInvalidType error.
// If the entity implements the interface, it checks if the ID is an ObjectID.
// If it is, it returns the ObjectID. Otherwise, it attempts to convert
// the ID to an ObjectID from its hexadecimal representation.
// If ID is nil returns a NilObjectID and a ErrInvalidHex error.
// Returns the ObjectID and an error, if any.
func getObjectID(entity any) (primitive.ObjectID, error) {
	v, ok := any(entity).(core.Entity)
	if !ok {
		return primitive.ObjectID{}, ErrInvalidType
	}

	if v.GetID() == nil {
		return NilObjectID, ErrInvalidHex
	}

	if id, ok := v.GetID().(primitive.ObjectID); ok {
		return id, nil
	}

	return primitive.ObjectIDFromHex(v.GetID().(string))
}
