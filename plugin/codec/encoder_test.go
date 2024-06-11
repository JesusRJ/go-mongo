package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ownerID, _ = primitive.ObjectIDFromHex("66396fb0465065406a8f3229")

type Product struct {
	SKU        string     `bson:"sku"`
	Price      float32    `bson:"price"`
	Owner      Owner      `ref:"belongsTo,owner,owner_id,_id"`
	Categories []Category `ref:"hasMany,category,category_id,_id,categories"`
}

type Owner struct {
	ID   primitive.ObjectID
	Name string
}

type Category struct {
	ID   primitive.ObjectID
	Name string
}

func (o Owner) GetID() any {
	return o.ID
}

type testCase struct {
	name string
	val  any
	want any
}

var testCases = []testCase{
	{
		name: "basic",
		val: Product{
			SKU:   "AB12345",
			Price: 399.0,
			Owner: Owner{
				ID:   ownerID,
				Name: "Smith Jr.",
			},
			Categories: []Category{},
		},
		want: struct {
			SKU        string             `bson:"sku"`
			Price      float32            `bson:"price"`
			Owner      primitive.ObjectID `bson:"owner_id"`
			Categories []Category         `bson:"-"`
		}{"AB12345", 399.0, ownerID, []Category{}},
	},
}

func TestEncoderEncode(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			enc, _ := NewEncoder()
			got, err := enc.Encode(tc.val)
			noerr(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
