package talkServer

import (
	"context"
	"fmt"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (t TalkHandler) RegisterPrimary(ctx context.Context, request *service.RegisterPrimaryRequest) (*service.RegisterPrimaryResponse, error) {
	jwt, err := authClient.VerifyIDToken(ctx, request.Token)
	if err != nil {
		return nil, status.New(codes.Unauthenticated, "authentication failed").Err()
	}
	md := metadata.New(map[string]string{"x-sakura-access": request.Token})
	err = grpc.SendHeader(ctx, md)
	grpc.SetTrailer(ctx, md)
	if err != nil {
		return nil, status.New(codes.Internal, "internal error").Err()
	}
	fmt.Println("register primary called")
	if _, err := dbClient.FetchUserAttribute(jwt.UID, bson.D{{"mid", 1}}); err != nil {
		err := dbClient.InsertNewUser(&talkDatabase.User{
			MID: jwt.UID,
		})
		if err != nil {
			return nil, status.New(codes.Internal, "internal error").Err()
		}
	}
	return &service.RegisterPrimaryResponse{}, nil
}
