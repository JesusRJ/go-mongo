package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestSaveWithPointerRef(t *testing.T) {
	tt := []struct {
		name   string
		input  *User
		fields []string // fields to validate
	}{
		{
			name:   "standalone",
			input:  GetUser("standalone", Config{}),
			fields: []string{"ID", "Name", "Address", "Pets"},
		},
		{
			name: "company_belongs_to",
			input: GetUser("company_belongs_to", Config{
				Company: &Company{
					Entity: db.Entity{ID: ObjectIDFromHex(StaticCompanyID[0])},
					Name:   "Teste",
				}}),
			fields: []string{"ID", "Name", "Address", "Pets"},
		},
	}

	repository, err := db.NewRepository[User](Database.Collection(CollUser))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := repository.Save(context.TODO(), tc.input)
			if err != nil {
				t.Fatalf("errors happened when create: %v", err)
			}

			if res.ID == nil || res.ID == "" {
				t.Errorf("user's primary key should has value after create, got : %v", tc.input.ID)
			}

			if res.CreatedAt.IsZero() {
				t.Errorf("user's created at should be not zero")
			}

			if tc.input.UpdatedAt.IsZero() {
				t.Errorf("user's updated at should be not zero")
			}

			newUser, err := repository.Find(context.Background(), tc.input)
			if err != nil {
				t.Fatalf("errors happened when find: %v", err)
			}

			AssertObjEqual(t, tc.input, newUser, tc.fields...)
		})
	}

}

func TestSaveWithoutPointerRef(t *testing.T) {
	tt := []struct {
		name   string
		input  Pet
		fields []string
	}{
		{
			name: "standalone",
			input: Pet{
				Name: "standalone",
			},
			fields: []string{"ID", "Name"},
		},
		{
			name: "user_belongs_to",
			input: Pet{
				User: *GetUser("dono", Config{ID: ObjectIDFromHex(StaticUserID[1])}),
				Name: "user_belongs_to",
			},
			fields: []string{"ID", "Name"},
		},
	}

	repository, err := db.NewRepository[Pet](Database.Collection(CollPet))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := repository.Save(context.TODO(), &tc.input)
			if err != nil {
				t.Fatalf("errors happened when create: %v", err)
			}

			if res.ID == nil || res.ID == "" {
				t.Errorf("user's primary key should has value after create, got : %v", tc.input.ID)
			}

			if res.CreatedAt.IsZero() {
				t.Errorf("user's created at should be not zero")
			}

			if tc.input.UpdatedAt.IsZero() {
				t.Errorf("user's updated at should be not zero")
			}

			newPet, err := repository.Find(context.Background(), &tc.input)
			if err != nil {
				t.Fatalf("errors happened when find: %v", err)
			}

			AssertObjEqual(t, tc.input, newPet, tc.fields...)
		})
	}
}

func TestCreateRegularlWithID(t *testing.T) {
	repository, err := db.NewRepository[RegularEntity](Database.Collection(CollAny))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	entity := RegularEntity{Name: "any name", Value: 10}
	res, err := repository.Save(context.TODO(), &entity)
	if err != nil {
		t.Fatalf("errors happened when create: %v", err)
	}

	if res.ID == "" {
		t.Errorf("primary key should has value after create, got : %v", entity.ID)
	}

	newUser, err := repository.Find(context.Background(), &entity)
	if err != nil {
		t.Fatalf("errors happened when find: %v", err)
	}

	AssertObjEqual(t, entity, newUser, "ID", "Name", "Value")
}

func TestCreateRegularlWithoutID(t *testing.T) {
	repository, err := db.NewRepository[RegularEntityWithoutID](Database.Collection(CollAny))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	entity := RegularEntityWithoutID{Name: "any name", Value: 10}
	res, err := repository.Save(context.TODO(), &entity)
	if err != nil {
		t.Fatalf("errors happened when create: %v", err)
	}

	if res.GetID() != "" {
		t.Errorf("primary key should has no value after create, got : %v", res.GetID())
	}
}

func TestSaveComplex(t *testing.T) {

}
