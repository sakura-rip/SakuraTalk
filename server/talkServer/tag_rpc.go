package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	err := dbClient.InsertOrUpdateUserTag(utils.GetUUID(ctx), dbTag)
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
	err := dbClient.InsertOrUpdateUserTag(utils.GetUUID(ctx), dbTag)
	if err != nil {
		return nil, err
	}
	return &service.UpdateTagResponse{}, nil
}

func (t TalkHandler) RegisterTags(ctx context.Context, request *service.RegisterTagsRequest) (*service.RegisterTagsResponse, error) {
	user, err := dbClient.FetchUserAttribute(request.Mid, bson.D{{"contacts", 1}, {"profile", 1}, {"tags", 1}})
	//付与先のユーザー存在確認
	if err != nil {
		return nil, err
	}
	//タグの存在確認
	_, ok := user.Tags[request.TagID]
	if !ok {
		return nil, status.New(codes.NotFound, "no such tag").Err()
	}
	contact, ok := user.Contacts[request.Mid]
	if !ok {
		contact = talkDatabase.Contact{
			MID:             request.Mid,
			OverWrittenName: user.Profile.Name,
		}
	}
	contact.TagIds = append(contact.TagIds, request.TagID)
	err = dbClient.UpdateUser(utils.GetUUID(ctx), bson.D{{"contacts." + request.Mid, contact}})
	if err != nil {
		return nil, err
	}
	return &service.RegisterTagsResponse{}, nil
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
