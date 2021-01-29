package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"
)

type TalkHandler struct{}

type DefaultMiddleWare func(ctx context.Context) (context.Context, error)

func newUnaryServerInterceptor(authFunc DefaultMiddleWare) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
		jwt, err := auth.VerifyIDToken(ctx, token)
		if err != nil {
			return ctx, status.New(codes.Unauthenticated, "authentication failed").Err()
		}
		ctx = utils.SetKeyAndValue(ctx, "uuid", jwt.UID)
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
	talkService.RegisterTalkServiceServer(server, ts)
	if err := server.Serve(listen); err != nil {
		panic(err)
	}
}
