package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) RegisterPrimary(ctx context.Context, request *service.RegisterPrimaryRequest) (*service.RegisterPrimaryResponse, error) {
	panic("implement me")
}
