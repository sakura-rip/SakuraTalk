package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
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

func (t TalkHandler) GetAllTags(ctx2 context.Context, empty *service.Empty) (*service.GetAllTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteTag(ctx context.Context, request *service.DeleteTagRequest) (*service.DeleteTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetTag(ctx context.Context, request *service.GetTagRequest) (*service.GetTagResponse, error) {
	panic("implement me")
}
