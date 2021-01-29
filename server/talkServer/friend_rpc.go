package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) AddFriend(ctx context.Context, request *talkService.AddFriendRequest) (*talkService.AddFriendResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteFriends(ctx context.Context, request *talkService.DeleteFriendsRequest) (*talkService.DeleteFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) BlockFriends(ctx context.Context, request *talkService.BlockFriendsRequest) (*talkService.BlockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnblockFriends(ctx context.Context, request *talkService.UnblockFriendsRequest) (*talkService.UnblockFriendsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AddFriendsToFavorite(ctx context.Context, request *talkService.AddFriendsToFavoriteRequest) (*talkService.AddFriendsToFavoriteResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveFriendsFromFavorite(ctx context.Context, request *talkService.RemoveFriendsFromFavoriteRequest) (*talkService.RemoveFriendsFromFavoriteResponse, error) {
	panic("implement me")
}
