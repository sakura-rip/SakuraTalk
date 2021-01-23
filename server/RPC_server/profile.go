package RPC_server

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) UpdateProfileAttributes(ctx context.Context, request *TalkService.UpdateProfileAttributesRequest) (*TalkService.UpdateProfileAttributesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetProfile(ctx context.Context, empty *TalkService.Empty) (*TalkService.Profile, error) {
	panic("implement me")
}
