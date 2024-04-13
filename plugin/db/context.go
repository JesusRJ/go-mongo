package db

import "context"

type contextKeyTx struct{}

func OfContext(ctx context.Context) {
	ctx.Value(contextKeyTx{})
}

func WithContext(ctx context.Context, value any) context.Context {
	return context.WithValue(ctx, contextKeyTx{}, value)
}
