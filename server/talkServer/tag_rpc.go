package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) CreateTag(ctx context.Context, request *talkService.CreateTagRequest) (*talkService.CreateTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateTag(ctx context.Context, request *talkService.UpdateTagRequest) (*talkService.UpdateTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RegisterTags(ctx context.Context, request *talkService.RegisterTagsRequest) (*talkService.RegisterTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetAllTags(ctx2 context.Context, empty *talkService.Empty) (*talkService.GetAllTagsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) DeleteTag(ctx context.Context, request *talkService.DeleteTagRequest) (*talkService.DeleteTagResponse, error) {
	panic("implement me")
}

func (t TalkHandler) GetTag(ctx context.Context, request *talkService.GetTagRequest) (*talkService.GetTagResponse, error) {
	panic("implement me")
}
