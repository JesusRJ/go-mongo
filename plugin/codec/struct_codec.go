package codec

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

type StructCodec struct {
	cache  sync.Map // map[reflect.Type]*structDescription
	parser StructTagParser
	x      bson.Decoder
}
