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
	//TODO: すでに同じ名前のTAGがないか
	tagID := utils.GenerateUUID()
	dbTag := talkDatabase.Tag{
		TagID:       tagID,
		Name:        request.Tag.Name,
		Description: request.Tag.Description,
		Color:       request.Tag.Color,
		Creator:     request.Tag.Creator,
		CreatedTime: request.Tag.CreatedTime,
	}
	err := dbClient.InsertUserTag(utils.GetMid(ctx), dbTag)
	if err != nil {
		return &service.CreateTagResponse{}, err
	}
	return &service.CreateTagResponse{TagID: tagID}, nil
}

func (t TalkHandler) UpdateTag(ctx context.Context, request *service.UpdateTagRequest) (*service.UpdateTagResponse, error) {
	//TODO: すでに同じ名前のTAGがないか
	dbTag := talkDatabase.Tag{
		TagID:       request.Tag.TagID,
		Name:        request.Tag.Name,
		Description: request.Tag.Description,
		Color:       request.Tag.Color,
		Creator:     request.Tag.Creator,
		CreatedTime: request.Tag.CreatedTime,
	}
	err := dbClient.InsertUserTag(utils.GetMid(ctx), dbTag)
	if err != nil {
		return &service.UpdateTagResponse{}, err
	}
	return &service.UpdateTagResponse{}, nil
}

func (t TalkHandler) RegisterTags(ctx context.Context, request *service.RegisterTagsRequest) (*service.RegisterTagsResponse, error) {
	user, err := dbClient.FetchUserAttributes(request.Mid, "contacts", "profile", "tags")
	//付与先のユーザー存在確認
	if err != nil {
		return &service.RegisterTagsResponse{}, err
	}
	//タグの存在確認
	_, ok := user.Tags[request.TagID]
	if !ok {
		return &service.RegisterTagsResponse{}, status.Error(codes.NotFound, "no such tag")
	}
	contact, ok := user.Contacts[request.Mid]
	if !ok {
		contact = talkDatabase.Contact{
			MID:             request.Mid,
			OverWrittenName: user.Profile.Name,
		}
	}
	contact.TagIds = append(contact.TagIds, request.TagID)
	err = dbClient.UpdateUser(utils.GetMid(ctx), bson.D{{"contacts." + request.Mid, contact}})
	if err != nil {
		return &service.RegisterTagsResponse{}, err
	}
	return &service.RegisterTagsResponse{}, nil
}

func (t TalkHandler) GetAllTags(ctx context.Context, empty *service.Empty) (*service.GetAllTagsResponse, error) {
	var tagIds []string
	user, err := dbClient.FetchUserAttributes(utils.GetMid(ctx), "tags")
	if err != nil {
		return &service.GetAllTagsResponse{}, err
	}
	for tagID := range user.Tags {
		tagIds = append(tagIds, tagID)
	}
	response := &service.GetAllTagsResponse{TagIds: tagIds}
	return response, err
}

func (t TalkHandler) DeleteTag(ctx context.Context, request *service.DeleteTagRequest) (*service.DeleteTagResponse, error) {
	err := dbClient.DeleteUserTag(utils.GetMid(ctx), request.TagId)
	return &service.DeleteTagResponse{}, err
}

func (t TalkHandler) GetTag(ctx context.Context, request *service.GetTagRequest) (*service.GetTagResponse, error) {
	tag, err := dbClient.FetchUserTag(utils.GetMid(ctx), request.TagId)
	if err != nil {
		return &service.GetTagResponse{}, err
	}
	response := &service.GetTagResponse{}
	response.Tag = tag.ConvertToRPCStruct()
	return response, err
}
