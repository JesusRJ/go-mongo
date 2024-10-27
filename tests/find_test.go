package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/plugin/db"
	. "github.com/jesusrj/go-mongo/utils/tests"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindByID(t *testing.T) {
	tt := []struct {
		name    string
		input   *User
		want    string
		wantErr bool
	}{
		{
			name:    "without_id",
			input:   GetUser("create", Config{}),
			want:    "000000000000000000000000",
			wantErr: true,
		},
		{
			name:  "with_id",
			input: GetUser("create", Config{ID: StaticUserID[0]}),
			want:  "661f17bffc35c18b2f85e975",
		},
	}

	repository, err := db.NewRepository[User](Database.Collection(CollUser))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := repository.FindByID(context.TODO(), tc.input)

			if err == nil && tc.wantErr {
				t.Fatalf("a error is expected when run FindByID")
			}

			if err != nil && !tc.wantErr {
				t.Fatalf("errors happened when run FindByID: %v", err)
			}

			if got != nil && got.ID == nil {
				t.Errorf("user's primary key should has value after FindByID, got : %v", tc.input.ID)
			}

			if got != nil {
				v, _ := got.ID.(primitive.ObjectID)
				want, _ := primitive.ObjectIDFromHex(tc.want)
				if v != want {
					t.Errorf("user's primary key is wrong, want: %s, got : %v", tc.want, tc.input.ID)
				}
			}
		})
	}

}
