package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetSetting(ctx context.Context, empty *service.Empty) (*service.Setting, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateSettingAttributes(ctx context.Context, request *service.UpdateSettingAttributesRequest) (*service.UpdateSettingAttributesResponse, error) {
	panic("implement me")
}
