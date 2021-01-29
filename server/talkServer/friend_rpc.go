package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) AddFriend(ctx context.Context, request *service.AddFriendRequest) (*service.AddFriendResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteFriends(ctx context.Context, request *service.DeleteFriendsRequest) (*service.DeleteFriendsResponse, error) {
	panic("implement me")
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
