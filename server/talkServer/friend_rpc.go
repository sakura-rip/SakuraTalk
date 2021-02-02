package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t TalkHandler) AddFriend(ctx context.Context, request *service.AddFriendRequest) (*service.AddFriendResponse, error) {
	mid := utils.GetMid(ctx)
	blockedMids, err := dbClient.FetchUserBlockedIds(mid)
	if err != nil {
		return nil, err
	}
	if utils.IsStrInSlice(blockedMids, request.Mid) {
		return nil, status.Error(codes.InvalidArgument, "request mid is blocked")
	}
	err = dbClient.AddToSetUserAttribute(mid, "friendIds", request.Mid)
	if err != nil {
		return nil, err
	}
	err = dbClient.AddToSetUserAttribute(request.Mid, "recvFriendIds", mid)
	if err != nil {
		return nil, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, 1)
	if err != nil {
		return nil, err
	}
	return &service.AddFriendResponse{}, err
}

func (t TalkHandler) DeleteFriends(ctx context.Context, request *service.DeleteFriendsRequest) (*service.DeleteFriendsResponse, error) {
	mid := utils.GetMid(ctx)
	friendMids, err := dbClient.FetchUserFriendIds(mid)
	if err != nil {
		return nil, err
	}
	if !utils.IsStrInSlice(friendMids, request.Mid) {
		return nil, status.Error(codes.InvalidArgument, "request mid is not friend")
	}
	err = dbClient.AddToSetUserAttribute(mid, "deletedIds", request.Mid)
	if err != nil {
		return nil, err
	}
	err = dbClient.AddToSetUserAttribute(request.Mid, "recvDeletedIds", mid)
	if err != nil {
		return nil, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, 3)
	if err != nil {
		return nil, err
	}
	return &service.DeleteFriendsResponse{}, nil
}

func (t TalkHandler) BlockFriends(ctx context.Context, request *service.BlockFriendsRequest) (*service.BlockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnblockFriends(ctx context.Context, request *service.UnblockFriendsRequest) (*service.UnblockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AddFriendsToFavorite(ctx context.Context, request *service.AddFriendsToFavoriteRequest) (*service.AddFriendsToFavoriteResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveFriendsFromFavorite(ctx context.Context, request *service.RemoveFriendsFromFavoriteRequest) (*service.RemoveFriendsFromFavoriteResponse, error) {
	panic("implement me")
}
