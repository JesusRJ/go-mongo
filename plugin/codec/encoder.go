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

// tNilObjectID represents the reflect.Type of the primitive.NilObjectID.
var tNilObjectID = reflect.TypeOf(primitive.NilObjectID)

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
		vField := rval.Field(i)
		sField := rval.Type().Field(i)

		structTag, err := e.parser.ParseStructTag(sField)
		if err != nil {
			return nil, nil, err
		}

		value, field := e.processField(vField, sField, structTag)
		fields = append(fields, *field)
		values = append(values, value)
	}

	return fields, values, nil
}

func (e *Encoder) processField(val reflect.Value, sf reflect.StructField, tag StructTag) (any, *reflect.StructField) {
	value := val.Interface()

	switch tag.Relation {
	case BelongsTo:
		// Convert field to primitive.ObjectID
		sf = reflect.StructField{
			Name: sf.Name,
			Type: tNilObjectID,
			Tag:  reflect.StructTag(fmt.Sprintf(`bson:"%s"`, tag.LocalField)),
		}
		if entity, ok := val.Interface().(core.Entity); ok {
			value = entity.GetID()
		}
	case HasMany:
		// Ignore fields like HasMany
		sf = reflect.StructField{
			Name: sf.Name,
			Type: sf.Type,
			Tag:  reflect.StructTag(`bson:"-"`),
		}
	}

	return value, &sf
}

func (e *Encoder) createStruct(fields []reflect.StructField, values []any) any {
	dType := reflect.StructOf(fields)
	newStruct := reflect.New(dType).Elem()

	for i, val := range values {
		newStruct.Field(i).Set(reflect.ValueOf(val))
	}

	return newStruct.Interface()
}
