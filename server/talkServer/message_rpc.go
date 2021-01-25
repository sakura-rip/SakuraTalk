package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) SendMessage(ctx context.Context, request *TalkService.SendMessageRequest) (*TalkService.SendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) EditMessage(ctx context.Context, request *TalkService.EditMessageRequest) (*TalkService.EditMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ResendMessage(ctx context.Context, request *TalkService.ResendMessageRequest) (*TalkService.ResendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnsendMessage(ctx context.Context, request *TalkService.UnsendMessageRequest) (*TalkService.UnsendMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReadMessage(ctx context.Context, request *TalkService.ReadMessageRequest) (*TalkService.ReadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UnreadMessage(ctx context.Context, request *TalkService.UnreadMessageRequest) (*TalkService.UnreadMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ReportMessage(ctx context.Context, request *TalkService.ReportMessageRequest) (*TalkService.ReportMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) RemoveAllMessages(ctx context.Context, request *TalkService.RemoveAllMessagesRequest) (*TalkService.RemoveAllMessagesResponse, error) {
	panic("implement me")
}

func (t TalkHandler) ActionMessage(ctx context.Context, request *TalkService.ActionMessageRequest) (*TalkService.ActionMessageResponse, error) {
	panic("implement me")
}

func (t TalkHandler) AnnounceMessage(ctx context.Context, request *TalkService.AnnounceMessageRequest) (*TalkService.AnnounceMessageResponse, error) {
	panic("implement me")
}
