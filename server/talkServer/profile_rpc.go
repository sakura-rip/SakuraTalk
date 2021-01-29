package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) UpdateProfileAttributes(ctx context.Context, request *talkService.UpdateProfileAttributesRequest) (*talkService.UpdateProfileAttributesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetProfile(ctx context.Context, empty *talkService.Empty) (*talkService.Profile, error) {
	panic("implement me")
}
