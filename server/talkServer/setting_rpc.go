package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (t TalkHandler) GetSetting(ctx context.Context, empty *service.Empty) (*service.Setting, error) {
	dbSetting, err := dbClient.FetchUserSetting(utils.GetMid(ctx))
	if err != nil {
		return &service.Setting{}, err
	}
	return dbSetting.ConvertToRPCStruct(), err
}

func (t TalkHandler) UpdateSettingAttributes(ctx context.Context, request *service.UpdateSettingAttributesRequest) (*service.UpdateSettingAttributesResponse, error) {
	var updateObject bson.D
	for _, key := range request.Key {
		switch key {
		case service.SettingKey_PRIVATE_USER_ID:
			updateObject = append(updateObject, bson.E{Key: "setting.pUserID", Value: request.Setting.PrivateUserID})
		case service.SettingKey_EMAIL:
			updateObject = append(updateObject, bson.E{Key: "setting.email", Value: request.Setting.Email})
		case service.SettingKey_ALLOW_SEARCH_BY_PRIVATE_USER_ID:
			updateObject = append(updateObject, bson.E{Key: "setting.asByPUserID", Value: request.Setting.AllowSearchByPrivateUserID})
		case service.SettingKey_ALLOW_SEARCH_BY_EMAIL:
			updateObject = append(updateObject, bson.E{Key: "setting.asByEmail", Value: request.Setting.AllowSearchByEmail})
		case service.SettingKey_ALLOW_SEARCH_BY_USER_TICKET:
			updateObject = append(updateObject, bson.E{Key: "setting.asByUserTicket", Value: request.Setting.AllowSearchByUserTicket})
		}
	}
	err := dbClient.UpdateUser(utils.GetMid(ctx), updateObject)
	if err != nil {
		return &service.UpdateSettingAttributesResponse{}, err
	}
	return &service.UpdateSettingAttributesResponse{}, nil
}

func (t TalkHandler) IssueUserTicket(ctx context.Context, request *service.IssueUserTicketRequest) (*service.IssueUserTicketResponse, error) {
	ticket := utils.GenerateUUID()
	err := dbClient.UpdateUser(utils.GetMid(ctx), bson.D{{"setting.UTicket", ticket}})
	if err != nil {
		return nil, err
	}
	return &service.IssueUserTicketResponse{Ticket: ticket}, nil
}
