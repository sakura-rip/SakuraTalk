package talkServer

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"google.golang.org/api/option"
	"google.golang.org/grpc/metadata"
)

var auth *firebaseAuth.Client
var ctx context.Context

// init firebase authentication
func init() {
	ctx = context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(fmt.Errorf("error on init firebase auth: %v", err))
	}
	auth, err = app.Auth(ctx)
	if err != nil {
		panic(fmt.Errorf("error on get auth %v", err))
	}
}

//getAccessToken get access Token from context
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
		return "", nil, status.New(codes.Unauthenticated, "authentication failed").Err()
	}
	jwt, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", nil, status.New(codes.Unauthenticated, "authentication failed").Err()
	}
	return jwt.UID, jwt.Claims, nil
}
