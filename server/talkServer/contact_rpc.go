package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetContacts(ctx context.Context, request *talkService.GetContactsRequest) (*talkService.GetContactsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateContact(ctx context.Context, request *talkService.UpdateContactRequest) (*talkService.UpdateContactResponse, error) {
	panic("implement me")
}
