package db

// import (
// 	"context"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type contextKeyTx struct{}

// func OfContext(ctx context.Context) (*mongo.Collection, bool) {
// 	coll, ok := ctx.Value(contextKeyTx{}).(*mongo.Collection)
// 	return coll, ok
// }

// func WithContext(ctx context.Context, value *mongo.Collection) context.Context {
// 	return context.WithValue(ctx, contextKeyTx{}, value)
// }
