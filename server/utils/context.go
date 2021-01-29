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

func GetUUID(ctx context.Context) string {
	return GetValue(ctx, "uuid").(string)
}

func GetClaims(ctx context.Context) map[string]interface{} {
	return GetValue(ctx, "claims").(map[string]interface{})
}
