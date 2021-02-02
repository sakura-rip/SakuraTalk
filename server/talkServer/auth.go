package talkServer

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// getHeader get access Token from context
func getHeader(ctx context.Context, key string) (string, bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	token := headers.Get(key)
	if len(token) == 0 {
		return "", false
	}
	return token[0], true
}

//VerifyTokenAndGetUUID check token and get token's user id
func VerifyTokenAndGetUUID(ctx context.Context) (string, map[string]interface{}, error) {
	token, ok := getHeader(ctx, "X-Sakura-Access")
	if !ok {
		return "", nil, status.Error(codes.Unauthenticated, "authentication failed")
	}
	jwt, err := authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", nil, status.Error(codes.Unauthenticated, "authentication failed")
	}
	return jwt.UID, jwt.Claims, nil
}
