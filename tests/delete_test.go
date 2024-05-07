package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"

	. "github.com/jesusrj/go-mongo/utils/tests"
)

func TestDelete(t *testing.T) {
	repository := db.NewRepository[User](Database.Collection(CollUser))

	tt := []struct {
		name  string
		input *User
		want  *User
	}{
		{
			name:  "Success Primitive ID",
			input: GetUser("update", Config{ID: ObjectIDFromHex(StaticUserID[3])}),
			want:  GetUser("update", Config{ID: ObjectIDFromHex(StaticUserID[3])}),
		},
		{
			name:  "Success String ID",
			input: GetUser("update", Config{ID: StaticUserID[4]}),
			want:  GetUser("update", Config{ID: StaticUserID[4]}),
		},
	}

	for _, tc := range tt {
		got, err := repository.Delete(context.TODO(), tc.input)
		if err != nil {
			t.Fatalf("test %s: expected: %v, got erro: %v", tc.name, tc.want, err)
		}

		AssertObjEqual(t, got, tc.want, "ID", "Name", "Address", "Pets")
	}
}
