package talkServer

import (
	"context"
	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"fmt"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
	"os"
)

var authClient *firebaseAuth.Client
var dbClient *talkDatabase.DBClient

// init firebase authentication
func init() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic(fmt.Errorf("error on init firebase authClient: %v", err))
	}
	authClient, err = app.Auth(ctx)
	if err != nil {
		panic(fmt.Errorf("error on get authClient %v", err))
	}
	dbClient = talkDatabase.NewDBClient()
}

type TalkHandler struct{}

type DefaultMiddleWare func(ctx context.Context) (context.Context, error)

func isPublicApi(methodName string) bool {
	switch methodName {
	case "/TalkService.TalkService/RegisterPrimary":
		return true
	default:
		return false
	}

}
func newUnaryServerInterceptor(authFunc DefaultMiddleWare) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		if isPublicApi(info.FullMethod) {
			return handler(ctx, req)
		}
		newCtx, err := authFunc(ctx)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, "authentication failed").Err()
		}
		return handler(newCtx, req)
	}
}

func newDefaultMiddleWare() DefaultMiddleWare {
	return func(ctx context.Context) (context.Context, error) {
		token, ok := getHeader(ctx, "X-Sakura-Access")
		if !ok {
			return ctx, status.New(codes.Unauthenticated, "authentication failed").Err()
		}
		jwt, err := authClient.VerifyIDToken(ctx, token)
		if err != nil {
			return ctx, status.New(codes.Unauthenticated, "authentication failed").Err()
		}
		ctx = utils.SetKeyAndValue(ctx, "mid", jwt.UID)
		ctx = utils.SetKeyAndValue(ctx, "claims", jwt.Claims)
		return ctx, nil
	}
}

// RunServer サーバーを走らせます！
func RunServer() {
	listen, err := net.Listen("tcp", ":8806")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(newUnaryServerInterceptor(newDefaultMiddleWare())),
	)
	ts := TalkHandler{}
	service.RegisterTalkServiceServer(server, ts)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
