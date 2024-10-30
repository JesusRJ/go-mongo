package db

import (
	"reflect"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson"
)

// filterWithID returns a MongoDB filter that targets a specific document by its ID,
// using the BSON format (bson.M{{"_id", ...}})
func filterWithID[T core.Entity](entity T) (bson.M, error) {
	id, err := getObjectID(entity)
	if err != nil {
		return nil, err
	}
	return bson.M{"_id": id}, nil
}

func filterWithFields[T core.Entity](entity T) bson.M {
	result := make(map[string]any)

	rval := reflect.ValueOf(entity)
	rtype := reflect.TypeOf(entity)

	if rtype.Kind() == reflect.Struct {
		for i := 0; i < rtype.NumField(); i++ {
			value := rval.Field(i)
			if value.Interface() == nil {
				continue
			}

			field := rtype.Field(i)

			bsonTag := field.Tag.Get("bson")
			if bsonTag == "" || bsonTag == "-" {
				bsonTag = field.Name
			}

			result[bsonTag] = value.Interface()
		}
	}

	return bson.M(result)
}
