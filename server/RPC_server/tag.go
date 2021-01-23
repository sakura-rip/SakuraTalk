package RPC_server

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) CreateTag(ctx context.Context, request *TalkService.CreateTagRequest) (*TalkService.CreateTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateTag(ctx context.Context, request *TalkService.UpdateTagRequest) (*TalkService.UpdateTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RegisterTags(ctx context.Context, request *TalkService.RegisterTagsRequest) (*TalkService.RegisterTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetAllTags(ctx context.Context, request *TalkService.GetAllTagsRequest) (*TalkService.GetAllTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteTag(ctx context.Context, request *TalkService.DeleteTagRequest) (*TalkService.DeleteTagResponse, error) {
	panic("implement me")
}
