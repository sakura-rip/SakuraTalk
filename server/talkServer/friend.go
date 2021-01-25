package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) AddFriend(ctx context.Context, request *TalkService.AddFriendRequest) (*TalkService.AddFriendResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteFriends(ctx context.Context, request *TalkService.DeleteFriendsRequest) (*TalkService.DeleteFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) BlockFriends(ctx context.Context, request *TalkService.BlockFriendsRequest) (*TalkService.BlockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnblockFriends(ctx context.Context, request *TalkService.UnblockFriendsRequest) (*TalkService.UnblockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AddFriendsToFavorite(ctx context.Context, request *TalkService.AddFriendsToFavoriteRequest) (*TalkService.AddFriendsToFavoriteResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveFriendsFromFavorite(ctx context.Context, request *TalkService.RemoveFriendsFromFavoriteRequest) (*TalkService.RemoveFriendsFromFavoriteResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateFriend(ctx context.Context, request *TalkService.UpdateFriendRequest) (*TalkService.UpdateFriendResponse, error) {
	panic("implement me")
}
