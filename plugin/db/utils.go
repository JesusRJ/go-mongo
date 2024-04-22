package db

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
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

	field.Set(reflect.ValueOf(value))
	return nil
}

// filterWithID returns a MongoDB filter that targets a specific document by its ID,
// using the BSON format (bson.M{{"_id", ...}})
func filterWithID[T core.AbstractEntity](entity T) (primitive.M, error) {
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

func toBsonM(e any) (bson.M, error) {
	buf := new(bytes.Buffer)
	vw, err := bsonrw.NewBSONValueWriter(buf)
	if err != nil {
		panic(err)
	}
	encoder, err := bson.NewEncoder(vw)
	if err != nil {
		panic(err)
	}

	registry := bson.NewRegistry()
	registry.RegisterTypeMapEntry(bson.TypeObjectID, reflect.TypeOf(nil))
	encoder.SetRegistry(registry)

	if err := encoder.Encode(e); err != nil {
		return nil, err
	}

	fmt.Println(bson.Raw(buf.Bytes()).String())
	return nil, nil
}
