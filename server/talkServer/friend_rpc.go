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
		return &service.AddFriendResponse{}, err
	}
	if utils.IsStrInSlice(blockedMids, request.Mid) {
		return &service.AddFriendResponse{}, status.Error(codes.InvalidArgument, "request mid is blocked")
	}
	err = dbClient.AddToSetUserAttribute(mid, "friendIds", request.Mid)
	if err != nil {
		return &service.AddFriendResponse{}, err
	}
	err = dbClient.AddToSetUserAttribute(request.Mid, "recvFriendIds", mid)
	if err != nil {
		return &service.AddFriendResponse{}, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, service.ContactStatus_FRIEND)
	if err != nil {
		return &service.AddFriendResponse{}, err
	}
	return &service.AddFriendResponse{}, err
}

func (t TalkHandler) DeleteFriends(ctx context.Context, request *service.DeleteFriendsRequest) (*service.DeleteFriendsResponse, error) {
	mid := utils.GetMid(ctx)
	friendMids, err := dbClient.FetchUserFriendIds(mid)
	if err != nil {
		return &service.DeleteFriendsResponse{}, err
	}
	if !utils.IsStrInSlice(friendMids, request.Mid) {
		return &service.DeleteFriendsResponse{}, status.Error(codes.InvalidArgument, "request mid is not friend")
	}
	err = dbClient.AddToSetUserAttribute(mid, "deletedIds", request.Mid)
	if err != nil {
		return &service.DeleteFriendsResponse{}, err
	}
	err = dbClient.AddToSetUserAttribute(request.Mid, "recvDeletedIds", mid)
	if err != nil {
		return &service.DeleteFriendsResponse{}, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, service.ContactStatus_DELETED)
	if err != nil {
		return &service.DeleteFriendsResponse{}, err
	}
	return &service.DeleteFriendsResponse{}, nil
}

func (t TalkHandler) BlockFriends(ctx context.Context, request *service.BlockFriendsRequest) (*service.BlockFriendsResponse, error) {
	mid := utils.GetMid(ctx)
	user, err := dbClient.FetchUserAttributes(mid, "friendIds", "blockedIds")
	if err != nil {
		return &service.BlockFriendsResponse{}, err
	}
	if utils.IsStrInSlice(user.FriendIds, request.Mid) {
		err = dbClient.RemoveFromSetUserAttribute(mid, "friendIds", request.Mid)
		if err != nil {
			return &service.BlockFriendsResponse{}, err
		}
	}
	//BLOCK済み
	if utils.IsStrInSlice(user.BlockedIds, request.Mid) {
		return &service.BlockFriendsResponse{}, status.Error(codes.InvalidArgument, "already blocked")
	}
	err = dbClient.AddToSetUserAttribute(mid, "blockedIds", request.Mid)
	if err != nil {
		return &service.BlockFriendsResponse{}, err
	}
	err = dbClient.AddToSetUserAttribute(request.Mid, "recvBlockedIds", mid)
	if err != nil {
		return &service.BlockFriendsResponse{}, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, service.ContactStatus_BLOCKED)
	if err != nil {
		return nil, err
	}
	return &service.BlockFriendsResponse{}, err
}

func (t TalkHandler) UnblockFriends(ctx context.Context, request *service.UnblockFriendsRequest) (*service.UnblockFriendsResponse, error) {
	mid := utils.GetMid(ctx)
	user, err := dbClient.FetchUserAttributes(mid, "blockedIds")
	if err != nil {
		return &service.UnblockFriendsResponse{}, err
	}
	if !utils.IsStrInSlice(user.BlockedIds, request.Mid) {
		return &service.UnblockFriendsResponse{}, status.Error(codes.InvalidArgument, "request mid is not blocked")
	}
	err = dbClient.RemoveFromSetUserAttribute(mid, "blockedIds", request.Mid)
	if err != nil {
		return &service.UnblockFriendsResponse{}, err
	}
	err = dbClient.RemoveFromSetUserAttribute(request.Mid, "recvBlockedIds", mid)
	if err != nil {
		return &service.UnblockFriendsResponse{}, err
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, service.ContactStatus_NO_RELATION)
	if err != nil {
		return &service.UnblockFriendsResponse{}, err
	}
	return &service.UnblockFriendsResponse{}, nil
}

func (t TalkHandler) AddFriendsToFavorite(ctx context.Context, request *service.AddFriendsToFavoriteRequest) (*service.AddFriendsToFavoriteResponse, error) {
	mid := utils.GetMid(ctx)
	user, err := dbClient.FetchUserAttributes(mid, "friendIds", "blockedIds")
	if err != nil {
		return nil, err
	}
	if !utils.IsStrInSlice(user.FriendIds, request.Mid) {
		return &service.AddFriendsToFavoriteResponse{}, status.Error(codes.InvalidArgument, "request mid is not friend")
	}
	if utils.IsStrInSlice(user.BlockedIds, request.Mid) {
		return &service.AddFriendsToFavoriteResponse{}, status.Error(codes.InvalidArgument, "request mid is blocked")
	}
	err = dbClient.UpdateUserContactStatus(mid, request.Mid, service.ContactStatus_FAVORITE)
	if err != nil {
		return nil, err
	}
	return &service.AddFriendsToFavoriteResponse{}, err
}

func (t TalkHandler) RemoveFriendsFromFavorite(ctx context.Context, request *service.RemoveFriendsFromFavoriteRequest) (*service.RemoveFriendsFromFavoriteResponse, error) {
	mid := utils.GetMid(ctx)
	user, err := dbClient.FetchUserAttributes(mid, "friendIds", "blockedIds")
	if err != nil {
		return nil, err
	}
	if !utils.IsStrInSlice(user.FriendIds, request.Mid) {
		return &service.RemoveFriendsFromFavoriteResponse{}, status.Error(codes.InvalidArgument, "request mid is not friend")
	}
	if utils.IsStrInSlice(user.BlockedIds, request.Mid) {
		return &service.RemoveFriendsFromFavoriteResponse{}, status.Error(codes.InvalidArgument, "request mid is blocked")
	}
	err = dbClient.UpdateUserContactIsFavorite(mid, request.Mid, false)
	if err != nil {
		return &service.RemoveFriendsFromFavoriteResponse{}, err
	}
	return &service.RemoveFriendsFromFavoriteResponse{}, nil

}
