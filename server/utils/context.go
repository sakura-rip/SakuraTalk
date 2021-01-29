package utils

import (
	"context"
)

func SetKeyAndValue(parents context.Context, key, value string) context.Context {
	return context.WithValue(parents, key, value)
}

func GetValue(ctx context.Context, key string) (string, bool) {
	v := ctx.Value(key)

	value, err := v.(string)
	if !err {
		return "", false
	}
	return value, true
}
