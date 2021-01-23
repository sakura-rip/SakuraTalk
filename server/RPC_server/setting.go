package RPC_server

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) GetSetting(ctx context.Context, empty *TalkService.Empty) (*TalkService.Setting, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateSettingAttributes(ctx context.Context, request *TalkService.UpdateSettingAttributesRequest) (*TalkService.UpdateSettingAttributesResponse, error) {
	panic("implement me")
}
