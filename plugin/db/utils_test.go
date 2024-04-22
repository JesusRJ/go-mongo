package db

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSetField(t *testing.T) {
	tt := []struct {
		name      string
		field     string
		value     any
		target    any
		want      any
		wantError bool
	}{
		{
			name:      "Success",
			field:     "Name",
			value:     "New Name",
			target:    &struct{ Name string }{Name: "Old Name"},
			want:      &struct{ Name string }{Name: "New Name"},
			wantError: false,
		},
		{
			name:      "FieldNotFound",
			field:     "Pass",
			value:     "New Name",
			target:    &struct{ Name string }{Name: "Old Name"},
			want:      &struct{ Name string }{Name: "Old Name"},
			wantError: true,
		},
		{
			name:      "UpdateStringWithObjectID",
			field:     "ID",
			value:     objectID("661f17bffc35c18b2f85e975"),
			target:    &struct{ ID string }{ID: ""},
			want:      &struct{ ID string }{ID: "661f17bffc35c18b2f85e975"},
			wantError: false,
		},
		{
			name:      "UpdateInterfaceWithObjectID",
			field:     "ID",
			value:     objectID("661f17bffc35c18b2f85e975"),
			target:    &struct{ ID interface{} }{ID: ""},
			want:      &struct{ ID interface{} }{ID: objectID("661f17bffc35c18b2f85e975")},
			wantError: false,
		},
	}

	for _, tc := range tt {
		err := setField(tc.target, tc.field, tc.value)
		if err != nil && !tc.wantError {
			t.Fatalf("test %s: expected: %v, got erro: %v", tc.name, tc.want, err)
		}

		if !reflect.DeepEqual(tc.want, tc.target) {
			t.Fatalf("test %s: expected: %v, got: %v", tc.name, tc.want, tc.target)
		}
	}
}

func TestFilterWithID(t *testing.T) {

	tt := []struct {
		name      string
		input     entity
		want      primitive.M
		wantError bool
	}{
		{
			name:  "Success",
			input: entity{objectID("661f22bf8a35841050e85503")},
			want:  primitive.M{"_id": objectID("661f22bf8a35841050e85503")},
		},
		{
			name:  "Success String",
			input: entity{"661f22bf8a35841050e85503"},
			want:  primitive.M{"_id": objectID("661f22bf8a35841050e85503")},
		},
	}

	for _, tc := range tt {
		got, err := filterWithID(tc.input)
		if err != nil && !tc.wantError {
			t.Fatalf("test %s: expected: %v, got erro: %v", tc.name, tc.want, err)
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("test %s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestToBsonM(t *testing.T) {
	toBsonM(struct {
		ID   primitive.ObjectID `bson:"_id,omitempty"`
		Name string             `bson:"name,omitempty"`
	}{
		ID:   objectID("661f22bf8a35841050e85503"),
		Name: "Teste toBson",
	})
}

func objectID(id string) primitive.ObjectID {
	obj, _ := primitive.ObjectIDFromHex(id)
	return obj
}

type entity struct {
	ID any
}

func (e entity) GetID() any {
	return e.ID
}
