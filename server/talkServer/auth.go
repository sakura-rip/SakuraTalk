package talkServer

import (
	"context"
	"fmt"
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
func getHeader(ctx context.Context, key string) (token string, ok bool) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	return headers.Get(key)[0], ok
}

//VerifyTokenAndGetUUID check token and get token's user id
func VerifyTokenAndGetUUID(ctx context.Context) (uuid string, claims map[string]interface{}) {
	token, ok := getHeader(ctx, "X-Sakura-Access")
	if !ok {
		return "", nil
	}
	jwt, err := auth.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", nil
	}
	return jwt.UID, jwt.Claims
}
