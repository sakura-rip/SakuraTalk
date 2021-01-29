package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) UpdateProfileAttributes(ctx context.Context, request *service.UpdateProfileAttributesRequest) (*service.UpdateProfileAttributesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetProfile(ctx context.Context, empty *service.Empty) (*service.Profile, error) {
	panic("implement me")
}
