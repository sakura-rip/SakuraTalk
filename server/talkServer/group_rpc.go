package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetGroup(ctx context.Context, request *service.GetGroupRequest) (*service.GetGroupResponse, error) {
	panic("implement me")
}
