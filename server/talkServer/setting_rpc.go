package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetSetting(ctx context.Context, empty *talkService.Empty) (*talkService.Setting, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateSettingAttributes(ctx context.Context, request *talkService.UpdateSettingAttributesRequest) (*talkService.UpdateSettingAttributesResponse, error) {
	panic("implement me")
}
