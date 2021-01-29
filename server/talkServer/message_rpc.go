package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) SendMessage(ctx context.Context, request *service.SendMessageRequest) (*service.SendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) EditMessage(ctx context.Context, request *service.EditMessageRequest) (*service.EditMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ResendMessage(ctx context.Context, request *service.ResendMessageRequest) (*service.ResendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnsendMessage(ctx context.Context, request *service.UnsendMessageRequest) (*service.UnsendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReadMessage(ctx context.Context, request *service.ReadMessageRequest) (*service.ReadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnreadMessage(ctx context.Context, request *service.UnreadMessageRequest) (*service.UnreadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReportMessage(ctx context.Context, request *service.ReportMessageRequest) (*service.ReportMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveAllMessages(ctx context.Context, request *service.RemoveAllMessagesRequest) (*service.RemoveAllMessagesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ActionMessage(ctx context.Context, request *service.ActionMessageRequest) (*service.ActionMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AnnounceMessage(ctx context.Context, request *service.AnnounceMessageRequest) (*service.AnnounceMessageResponse, error) {
	panic("implement me")
}
