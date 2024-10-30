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
var tNilObjectID = reflect.TypeOf(&primitive.NilObjectID)

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

	return e.createStruct(fields, values)
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
		// Create a new definition for the field, replacing it with ObjectID type and bson tag
		sf = reflect.StructField{
			Name: sf.Name,
			Type: tNilObjectID,
			Tag:  reflect.StructTag(fmt.Sprintf(`bson:"%s"`, tag.LocalField)),
		}

		// Converte o valor do campo para um ObjectID, se for uma entidade
		if entity, ok := value.(core.Entity); ok {
			value = convertEntityToObjectID(val, entity)
		}

	case HasMany:
		// Set the struct field to ignore the field in MongoDB using tag "bson:\"-\""
		sf = reflect.StructField{
			Name: sf.Name,
			Type: sf.Type,
			Tag:  reflect.StructTag(`bson:"-"`),
		}
	}

	return value, &sf
}

func (e *Encoder) createStruct(fields []reflect.StructField, values []any) (any, error) {
	structType := reflect.StructOf(fields)
	structValue := reflect.New(structType).Elem()

	for i, field := range fields {
		if values[i] == nil {
			continue
		}
		fieldValue := reflect.ValueOf(values[i])

		if field.Type != fieldValue.Type() {
			return nil, fmt.Errorf("incorrect type for field %s: expected %s but got %s",
				field.Name, field.Type, fieldValue.Type())
		}

		structValue.Field(i).Set(fieldValue)
	}

	return structValue.Interface(), nil
}

// Auxiliary function to convert an entity into an ObjectID
func convertEntityToObjectID(val reflect.Value, entity core.Entity) any {
	if val.Kind() == reflect.Pointer {
		if !val.IsNil() {
			if id := entity.GetID(); id != nil {
				if oid, ok := id.(primitive.ObjectID); ok {
					return &oid
				}
			}
		}
	} else {
		if id := entity.GetID(); id != nil {
			if oid, ok := id.(primitive.ObjectID); ok {
				return &oid
			}
		}
	}
	return nil
}
