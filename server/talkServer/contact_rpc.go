package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
)

func (t TalkHandler) GetContacts(ctx context.Context, request *service.GetContactsRequest) (*service.GetContactsResponse, error) {
	panic("implement me")
}

func (t TalkHandler) UpdateContact(ctx context.Context, request *service.UpdateContactRequest) (*service.UpdateContactResponse, error) {
	panic("implement me")
}
