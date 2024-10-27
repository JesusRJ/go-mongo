package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestUpdate(t *testing.T) {
	repository, err := db.NewRepository[User](Database.Collection(CollUser))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	tt := []struct {
		name  string
		input *User
		want  *User
	}{
		{
			name:  "Success Primitive ID",
			input: GetUser("update", Config{ID: ObjectIDFromHex(StaticUserID[1]), Address: &Address{Street: "Avocato"}, Pets: 4}),
			want:  GetUser("update", Config{ID: ObjectIDFromHex(StaticUserID[1]), Address: &Address{Street: "Avocato"}, Pets: 4}),
		},
		{
			name:  "Success String ID",
			input: GetUser("update", Config{ID: StaticUserID[2], Address: &Address{Street: "Avocato"}}),
			want:  GetUser("update", Config{ID: StaticUserID[2], Address: &Address{Street: "Avocato"}}),
		},
	}

	for _, tc := range tt {
		got, err := repository.Update(context.TODO(), tc.input)
		if err != nil {
			t.Fatalf("test %s: expected: %v, got erro: %v", tc.name, tc.want, err)
		}

		if got.UpdatedAt.IsZero() {
			t.Errorf("user's updated at should be not zero")
		}

		AssertObjEqual(t, got, tc.want, "ID", "Name", "Address", "Pets")
	}
}
