package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"time"
)

func (t TalkHandler) GetGroups(ctx context.Context, request *service.GetGroupsRequest) (*service.GetGroupsResponse, error) {
	mid := utils.GetMid(ctx)
	var groups = map[string]*service.Group{}
	for _, gid := range request.GroupIds {
		group, err := dbClient.FetchGroup(gid)
		if err != nil {
			continue
		}
		groupSetting, _ := dbClient.FetchUserGroupSettings(mid, gid)
		groups[gid] = &service.Group{
			GroupID:                   gid,
			GroupName:                 group.Name,
			OverWrittenName:           groupSetting.OverWrittenName,
			Creator:                   group.CreatorMID,
			CreatedTime:               group.CreatedTime,
			IconPath:                  group.IconPath,
			CoverPath:                 group.CoverPath,
			IsFavorite:                groupSetting.IsFavorite,
			EnableNotification:        groupSetting.EnableNotification,
			EnableNotificationMention: groupSetting.EnableNotificationMention,
			EnableNotificationOnJoin:  groupSetting.EnableNotificationOnJoin,
			EnableNotificationOnKick:  groupSetting.EnableNotificationOnKick,
			TagIDs:                    groupSetting.TagIds,
			InvitationTicket:          group.GroupTicket,
			MemberIds:                 group.MemberMids,
			InvitedIds:                group.InvitedMids,
		}
	}
	return &service.GetGroupsResponse{Groups: groups}, nil
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
