package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestSave(t *testing.T) {
	user := *GetUser("create", Config{})
	repository := db.NewRepository[User](Database.Collection(CollUser))

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

func TestSaveComplex(t *testing.T) {

}

func TestCreateLiteral(t *testing.T) {

}
