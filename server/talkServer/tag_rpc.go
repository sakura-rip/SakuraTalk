package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (t TalkHandler) CreateTag(ctx context.Context, request *service.CreateTagRequest) (*service.CreateTagResponse, error) {
	tagID := utils.GenerateUUID()
	dbTag := talkDatabase.Tag{
		TagID:       tagID,
		Name:        request.Tag.Name,
		Description: request.Tag.Description,
		Color:       request.Tag.Color,
		Creator:     request.Tag.Creator,
		CreatedTime: request.Tag.CreatedTime,
	}
	err := dbClient.InsertUserTag(utils.GetUUID(ctx), dbTag)
	if err != nil {
		return nil, err
	}
	return &service.CreateTagResponse{TagID: tagID}, nil
}

func (t TalkHandler) UpdateTag(ctx context.Context, request *service.UpdateTagRequest) (*service.UpdateTagResponse, error) {
	dbTag := talkDatabase.Tag{
		TagID:       request.Tag.TagID,
		Name:        request.Tag.Name,
		Description: request.Tag.Description,
		Color:       request.Tag.Color,
		Creator:     request.Tag.Creator,
		CreatedTime: request.Tag.CreatedTime,
	}
	err := dbClient.InsertUserTag(utils.GetUUID(ctx), dbTag)
	if err != nil {
		return nil, err
	}
	return &service.UpdateTagResponse{}, nil
}

func (t TalkHandler) RegisterTags(ctx context.Context, request *service.RegisterTagsRequest) (*service.RegisterTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetAllTags(ctx context.Context, empty *service.Empty) (*service.GetAllTagsResponse, error) {
	var tagIds []string
	user, err := dbClient.FetchUserAttribute(utils.GetUUID(ctx), bson.D{{"tags", 1}})
	if err != nil {
		return nil, err
	}
	for tagID := range user.Tags {
		tagIds = append(tagIds, tagID)
	}
	response := &service.GetAllTagsResponse{TagIds: tagIds}
	return response, err
}

func (t TalkHandler) DeleteTag(ctx context.Context, request *service.DeleteTagRequest) (*service.DeleteTagResponse, error) {
	err := dbClient.DeleteUserTag(utils.GetUUID(ctx), request.TagId)
	return &service.DeleteTagResponse{}, err
}

func (t TalkHandler) GetTag(ctx context.Context, request *service.GetTagRequest) (*service.GetTagResponse, error) {
	tag, err := dbClient.FetchUserTag(utils.GetUUID(ctx), request.TagId)
	if err != nil {
		return nil, err
	}
	response := &service.GetTagResponse{}
	response.Tag = tag.ConvertToRPCStruct()
	return response, err
}
