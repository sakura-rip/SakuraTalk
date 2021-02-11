package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"time"
)

func (t TalkHandler) GetGroups(ctx context.Context, request *service.GetGroupsRequest) (*service.GetGroupsResponse, error) {
}

func (t TalkHandler) CreateGroup(ctx context.Context, request *service.CreateGroupRequest) (*service.CreateGroupResponse, error) {
	mid := utils.GetMid(ctx)
	newGid := utils.GenerateUUID()
	err := dbClient.InsertNewGroup(&talkDatabase.Group{
		GID:                    newGid,
		CreatorMID:             mid,
		CreatedTime:            time.Now().Unix(),
		MemberMids:             []string{mid},
		InvitedMids:            request.InviteMids,
		Name:                   request.Name,
		Description:            request.Description,
		IconPath:               request.IconPath,
		CoverPath:              "",
		GroupTicket:            "",
		AllowJoinByGroupTicket: false,
	})
	if err != nil {
		return &service.CreateGroupResponse{}, err
	}
	return &service.CreateGroupResponse{Gid: newGid}, err
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
