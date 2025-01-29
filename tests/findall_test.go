package tests_test

import (
	"context"
	"testing"

	"github.com/jesusrj/go-mongo/core"
	"github.com/jesusrj/go-mongo/plugin/db"
	. "github.com/jesusrj/go-mongo/utils/tests"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFindAll(t *testing.T) {
	tt := []struct {
		name   string
		opts   core.QueryOptions // required
		query  any
		fields []string // fields to validate
		want   core.Pagination[User]
	}{
		// {
		// 	name:   "without_query",
		// 	opts:   *core.Options(),
		// 	query:  bson.D{},
		// 	fields: []string{"Page", "Pages", "Total"},
		// 	want: core.Pagination[User]{
		// 		Page:  1,
		// 		Pages: 50,
		// 		Total: 505,
		// 	},
		// },
		{
			name:   "with_query",
			opts:   *core.Options(),
			query:  bson.D{{Key: "name", Value: "user_batch_4"}},
			fields: []string{"Page", "Pages", "Total"},
			want: core.Pagination[User]{
				Page:  1,
				Pages: 0,
				Total: 1,
			},
		},
		// order:    core.OrderBy{Field: "name", Direction: core.Asc},
	}

	repository, err := db.NewPaginatedRepository[User, any](Database.Collection(CollUser))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := repository.FindAll(context.TODO(), tc.query, tc.opts)
			if err != nil {
				t.Fatalf("errors happened when run FindAll: %v", err)
			}

			if len(got.Data) != int(*tc.opts.PageSize) {
				t.Errorf("pages total, got : %v, expected:  %v", len(got.Data), *tc.opts.PageSize)
			}

			AssertObjEqual(t, tc.want, got, tc.fields...)
		})
	}
}
