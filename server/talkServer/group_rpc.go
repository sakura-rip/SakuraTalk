package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetGroup(ctx context.Context, request *service.GetGroupRequest) (*service.GetGroupResponse, error) {
	panic("implement me")
}

func (t TalkHandler) CreateGroup(ctx context.Context, request *service.CreateGroupRequest) (*service.CreateGroupResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateGroup(ctx context.Context, request *service.UpdateGroupRequest) (*service.UpdateGroupResponse, error) {
	panic("implement me")
}

func (t TalkHandler) InviteIntoGroup(ctx context.Context, request *service.InviteIntoGroupRequest) (*service.InviteIntoGroupResponse, error) {
	panic("implement me")
}

func (t TalkHandler) KickoutFromGroup(ctx context.Context, request *service.KickoutFromGroupRequest) (*service.KickoutFromGroupResponse, error) {
	panic("implement me")
}
