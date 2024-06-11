package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ownerID, _ = primitive.ObjectIDFromHex("66396fb0465065406a8f3229")

type Product struct {
	SKU   string  `bson:"sku"`
	Owner Owner   `ref:"belongsTo,owner,owner_id,_id"`
	Price float32 `bson:"price"`
}

type Owner struct {
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
			SKU: "AB12345",
			Owner: Owner{
				ID:   ownerID,
				Name: "Smith Jr.",
			},
			Price: 399.0,
		},
		want: struct {
			SKU   string             `bson:"sku"`
			Owner primitive.ObjectID `bson:"owner_id"`
			Price float32            `bson:"price"`
		}{"AB12345", ownerID, 399.0},
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
