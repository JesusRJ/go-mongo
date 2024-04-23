package tests_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestUpdate(t *testing.T) {
	user := *GetUser("update", Config{ID: StaticID[1]})
	repository := db.NewRepository[User](Database.Collection(CollUser))

	user.Address = fmt.Sprintf("854 Avocado Ave. - Newport Beach (CA) %v", time.Now())

	updatedUser, err := repository.Update(context.TODO(), &user)
	if err != nil {
		t.Fatalf("errors happened when update: %v", err)
	}

	if user.UpdatedAt.IsZero() {
		t.Errorf("user's updated at should be not zero")
	}

	AssertObjEqual(t, user, updatedUser, "ID", "Name", "Address", "Pets")
}
