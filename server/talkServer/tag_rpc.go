package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t TalkHandler) CreateTag(ctx context.Context, request *service.CreateTagRequest) (*service.CreateTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateTag(ctx context.Context, request *service.UpdateTagRequest) (*service.UpdateTagResponse, error) {
	panic("implement me")
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
	_, err := dbClient.UserCol.UpdateOne(
		ctx, bson.M{"_id": utils.GetUUID(ctx)},
		bson.M{"$unset": "tags." + request.TagId})
	if err != nil {
		return nil, status.New(codes.NotFound, "tag not found").Err()
	}
	return &service.DeleteTagResponse{}, nil
}

func (t TalkHandler) GetTag(ctx context.Context, request *service.GetTagRequest) (*service.GetTagResponse, error) {
	tag, err := dbClient.FetchUserTag(utils.GetUUID(ctx), request.TagId)
	if err != nil {
		return nil, err
	}
	response := &service.GetTagResponse{}
	response.Tag = &service.Tag{
		TagID:       request.TagId,
		Name:        tag.Name,
		Description: tag.Description,
		Color:       tag.Color,
		Creator:     tag.Creator,
		CreatedTime: tag.CreatedTime,
	}
	return response, err
}
