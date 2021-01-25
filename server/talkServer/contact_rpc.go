package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/TalkService"
)

func (t TalkHandler) GetContacts(ctx context.Context, request *TalkService.GetContactsRequest) (*TalkService.GetContactsResponse, error) {
	panic("implement me")
}
