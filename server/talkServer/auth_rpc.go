package talkServer

import (
	"context"
	"fmt"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (t TalkHandler) RegisterPrimary(ctx context.Context, request *service.RegisterPrimaryRequest) (*service.RegisterPrimaryResponse, error) {
	_, err := authClient.VerifyIDToken(ctx, request.Token)
	fmt.Println(err)
	if err != nil {
		return nil, status.New(codes.Unauthenticated, "authentication failed").Err()
	}
	err = grpc.SetHeader(ctx, metadata.Pairs("x-sakura-access", request.Token))
	fmt.Println(err)
	if err != nil {
		return nil, status.New(codes.Internal, "internal error").Err()
	}
	return &service.RegisterPrimaryResponse{}, nil
}
