package utils

import (
	"context"
)

func SetKeyAndValue(parents context.Context, key string, value interface{}) context.Context {
	return context.WithValue(parents, key, value)
}

func GetValue(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}
