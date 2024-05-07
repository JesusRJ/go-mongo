package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"
	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestFind(t *testing.T) {
	user := *GetUser("create", Config{ID: StaticUserID[0]})
	repository := db.NewRepository[User](Database.Collection(CollUser))

	res, err := repository.Find(context.TODO(), &user)
	if err != nil {
		t.Fatalf("errors happened when find: %v", err)
	}

	if res.ID == "" {
		t.Errorf("user's primary key should has value after find, got : %v", user.ID)
	}

	if res.CreatedAt.IsZero() {
		t.Errorf("user's finded at should be not zero")
	}
}
