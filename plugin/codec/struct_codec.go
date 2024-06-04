package codec

import (
	"sync"

	"go.mongodb.org/mongo-driver/bson"
)

// see:
// bson/bsoncodec/struct_codec.go
// bson/decoder.go
// bson/unmarshal.go

type StructCodec struct {
	cache  sync.Map // map[reflect.Type]*structDescription
	parser StructTagParser
	x      bson.Decoder
}
