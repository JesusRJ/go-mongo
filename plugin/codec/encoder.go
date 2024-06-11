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

// TODO: AVALIAR SE OS CAMPOS SÃO EXPORTADOS E IGNORAR OS NÃO EXPORTADOS
func (e *Encoder) Encode(val any) (any, error) {
	rval := reflect.ValueOf(val)
	if rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}

	fields, values, err := e.extractFieldsAndValues(rval)
	if err != nil {
		return nil, err
	}

	return e.createStruct(fields, values), nil
}

func (e *Encoder) extractFieldsAndValues(rval reflect.Value) ([]reflect.StructField, []any, error) {
	var fields []reflect.StructField
	var values []any

	for i := 0; i < rval.NumField(); i++ {
		valueField := rval.Field(i)
		typeField := rval.Type().Field(i)

		structTag, err := e.parser.ParseStructTag(typeField)
		if err != nil {
			return nil, nil, err
		}

		value, updatedField := e.processField(valueField, typeField, structTag)
		if updatedField == nil {
			continue // Ignore fields like HasMany
		}

		fields = append(fields, *updatedField)
		values = append(values, value)
	}

	return fields, values, nil
}

func (e *Encoder) processField(valueField reflect.Value, typeField reflect.StructField, structTag StructTag) (any, *reflect.StructField) {
	value := valueField.Interface()

	switch structTag.Relation {
	case BelongsTo:
		// Convert field to primitive.ObjectID
		typeField = reflect.StructField{
			Name: typeField.Name,
			Type: reflect.TypeOf(primitive.NilObjectID),
			Tag:  reflect.StructTag(fmt.Sprintf(`bson:"%s"`, structTag.LocalField)),
		}
		if entity, ok := valueField.Interface().(core.Entity); ok {
			value = entity.GetID()
		}
	case HasMany:
		return nil, nil // Ignore fields like HasMany
	}

	return value, &typeField
}

func (e *Encoder) createStruct(fields []reflect.StructField, values []any) any {
	dType := reflect.StructOf(fields)
	newStruct := reflect.New(dType).Elem()

	for i, val := range values {
		newStruct.Field(i).Set(reflect.ValueOf(val))
	}

	return newStruct.Interface()
}
