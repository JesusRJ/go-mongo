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
	repository, err := db.NewPaginatedRepository[User, any](Database.Collection(CollUser))
	if err != nil {
		t.Fatalf("errors happened when create repository: %v", err)
	}

	// query := bson.D{
	// 	{Key: "name", Value: "pet_1"},
	// 	// {Key: "age", Value: 9},
	// }

	p, err := repository.FindAll(context.TODO(), bson.D{}, *core.Options())
	if err != nil {
		t.Fatalf("errors happened when run FindAll: %v", err)
	}

	if p.Total != 505 {
		t.Errorf("records total, got : %v, expected: 505", p.Total)
	}

	if len(p.Data) != 10 {
		t.Errorf("pages total, got : %v, expected: 10", len(p.Data))
	}

	if p.Pages != 50 {
		t.Errorf("pages total, got : %v, expected: 50", p.Pages)
	}

	if p.Page != 1 {
		t.Errorf("wrong page, got : %v, expected: 1", p.Page)
	}
}
