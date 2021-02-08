package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (t TalkHandler) RegisterPrimary(ctx context.Context, request *service.RegisterPrimaryRequest) (*service.RegisterPrimaryResponse, error) {
	jwt, err := authClient.VerifyIDToken(ctx, request.Token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "authentication failed")
	}
	md := metadata.New(map[string]string{"x-sakura-access": request.Token})
	err = grpc.SendHeader(ctx, md)
	grpc.SetTrailer(ctx, md)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	if _, err := dbClient.FetchUserAttributes(jwt.UID, "mid"); err != nil {
		err := dbClient.InsertNewUser(&talkDatabase.User{
			MID:                jwt.UID,
			JoinedGroupIds:     []string{},
			InvitedGroupIds:    []string{},
			FriendIds:          []string{},
			BlockedIds:         []string{},
			DeletedIds:         []string{},
			ReceivedFriendIds:  []string{},
			ReceivedBlockedIds: []string{},
			ReceivedDeletedIds: []string{},
			GroupSettings:      map[string]talkDatabase.GroupSetting{},
			Tags:               map[string]talkDatabase.Tag{},
			Contacts:           map[string]talkDatabase.Contact{},
		})
		if err != nil {
			return nil, status.Error(codes.Internal, "internal error")
		}
	}
	return &service.RegisterPrimaryResponse{}, nil
}
