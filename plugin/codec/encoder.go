package codec

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/jesusrj/go-mongo/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ErrEncodeNil is the error returned when trying to encode a nil value
var ErrEncodeNil = errors.New("cannot Encode nil value")

type Encoder struct {
	parser StructTagParser
}

func NewEncoder() (*Encoder, error) {
	return &Encoder{
		parser: DefaultStructTagParser,
	}, nil
}

// TODO: AVALIAR SE OS CAMPOS SÃO EXPORTADOS
func (e *Encoder) Encode(val any) (any, error) {
	rval := reflect.ValueOf(val)
	if rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}

	fields := []reflect.StructField{}
	values := []any{}

	for x := 0; x < rval.NumField(); x++ {
		valueField := rval.Field(x)
		typeField := rval.Type().Field(x)

		structTag, err := e.parser.ParseStructTag(typeField)
		if err != nil {
			return nil, err
		}

		value := valueField.Interface()

		switch structTag.Relation {
		case BelongsTo:
			// Convert field to primitive.ObjectID
			typeField = reflect.StructField{
				Name: typeField.Name,
				Type: reflect.TypeOf(primitive.NilObjectID),
				Tag:  reflect.StructTag(fmt.Sprintf(`bson:"%s"`, structTag.LocalField)),
			}
			value = valueField.Interface().(core.Entity).GetID()
		case HasMany:
			continue // Ignore hasMany fields
		}

		fields = append(fields, typeField)
		values = append(values, value)
	}

	dType := reflect.StructOf(fields)
	value := reflect.New(dType)

	for v := range values {
		value.Elem().Field(v).Set(reflect.ValueOf(values[v]))
	}

	return value.Interface(), nil
}
