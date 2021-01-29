package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) SendMessage(ctx context.Context, request *talkService.SendMessageRequest) (*talkService.SendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) EditMessage(ctx context.Context, request *talkService.EditMessageRequest) (*talkService.EditMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ResendMessage(ctx context.Context, request *talkService.ResendMessageRequest) (*talkService.ResendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnsendMessage(ctx context.Context, request *talkService.UnsendMessageRequest) (*talkService.UnsendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReadMessage(ctx context.Context, request *talkService.ReadMessageRequest) (*talkService.ReadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnreadMessage(ctx context.Context, request *talkService.UnreadMessageRequest) (*talkService.UnreadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReportMessage(ctx context.Context, request *talkService.ReportMessageRequest) (*talkService.ReportMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveAllMessages(ctx context.Context, request *talkService.RemoveAllMessagesRequest) (*talkService.RemoveAllMessagesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ActionMessage(ctx context.Context, request *talkService.ActionMessageRequest) (*talkService.ActionMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AnnounceMessage(ctx context.Context, request *talkService.AnnounceMessageRequest) (*talkService.AnnounceMessageResponse, error) {
	panic("implement me")
}
