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
	rval := reflect.ValueOf(entity)
	rtype := rval.Type()

	result := make(map[string]any)

	// Add ID to filter
	if v, ok := any(entity).(core.Entity); ok {
		if id := v.GetID(); id != nil {
			result["_id"] = v.GetID()
		}
	}

	// Add additional fields to filter
	if rtype.Kind() == reflect.Struct {
		for i := 0; i < rtype.NumField(); i++ {
			value := rval.Field(i)
			field := rtype.Field(i)

			// Skip fields that are nil (for pointers, interfaces)
			if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
				if value.IsNil() {
					continue
				}
			}

			// Skip empty slices, maps or arrays
			switch value.Kind() {
			case reflect.Slice, reflect.Map, reflect.Array:
				if value.Len() == 0 {
					continue
				}
			}

			// Get the bson tag or use the field name if no bson tag is specified
			bsonTag := field.Tag.Get("bson")
			if bsonTag == "" || bsonTag == "-" {
				bsonTag = field.Name
			}

			if bsonTag == "inline" {
				continue
			}

			// Add the field value to the result map
			result[bsonTag] = value.Interface()
		}
	}

	return bson.M(result)
}
