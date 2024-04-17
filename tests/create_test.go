package tests_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"
	"go.mongodb.org/mongo-driver/bson/primitive"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestSave(t *testing.T) {
	user := *GetUser("create", Config{})
	repository := db.NewRepository[User](Database.Collection(CollUser))

	res, err := repository.Save(context.TODO(), &user)
	if err != nil {
		t.Fatalf("errors happened when create: %v", err)
	}

	if res.ID == "" {
		t.Errorf("user's primary key should has value after create, got : %v", user.ID)
	}

	if _, err := primitive.ObjectIDFromHex(res.ID); err != nil {
		t.Errorf("user's primary key invalid after create, got : %v", user.ID)
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

	if !reflect.DeepEqual(user, newUser) {
		t.Error("should be equals")
	}
}

func TestSaveComplex(t *testing.T) {

}

func TestCreateLiteral(t *testing.T) {

}
