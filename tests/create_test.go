package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestSave(t *testing.T) {
	repository := db.NewRepository[User](Database.Collection(CollUser))

	user := *GetUser("create", Config{})
	res, err := repository.Save(context.TODO(), &user)
	if err != nil {
		t.Fatalf("errors happened when create: %v", err)
	}

	if res.ID == nil || res.ID == "" {
		t.Errorf("user's primary key should has value after create, got : %v", user.ID)
	}

	if res.CreatedAt.IsZero() {
		t.Errorf("user's created at should be not zero")
	}

	if user.UpdatedAt.IsZero() {
		t.Errorf("user's updated at should be not zero")
	}

	newUser, err := repository.Find(context.Background(), &user)
	if err != nil {
		t.Fatalf("errors happened when find: %v", err)
	}

	AssertObjEqual(t, user, newUser, "ID", "Name", "Address", "Pets")
}

func TestCreateRegularlWithID(t *testing.T) {
	repository := db.NewRepository[RegularEntity](Database.Collection(CollAny))

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
	repository := db.NewRepository[RegularEntityWithoutID](Database.Collection(CollAny))

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
