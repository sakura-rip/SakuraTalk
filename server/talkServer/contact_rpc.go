package talkServer

import (
	"context"
	"github.com/sakura-rip/SakuraTalk/talkDatabase"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
)

func (t TalkHandler) GetContacts(ctx context.Context, request *service.GetContactsRequest) (*service.GetContactsResponse, error) {
	mid := utils.GetMid(ctx)
	var contacts = map[string]*service.Contact{}
	for _, targetMid := range request.Mids {
		targetProfile, err := dbClient.FetchUserProfile(targetMid)
		if err != nil {
			continue
		}
		targetContact, err := dbClient.FetchUserContact(mid, targetMid)
		if err != nil {
			continue
		}
		contacts[targetMid] = &service.Contact{
			Mid:             targetMid,
			DisplayName:     targetProfile.Name,
			OverWrittenName: targetContact.OverWrittenName,
			Bio:             targetProfile.Bio,
			IconPath:        targetProfile.IconPath,
			CoverPath:       targetProfile.CoverPath,
			TwitterID:       targetProfile.TwitterID,
			InstagramID:     targetProfile.InstagramID,
			GithubID:        targetProfile.GithubID,
			WebSiteURL:      targetProfile.WebSiteURL,
			Location:        targetProfile.Location,
			BirthDay:        targetProfile.BirthDay,
			Status:          targetProfile.Status,
			ContactStatus:   talkDatabase.RPCContactStatusFromDB(targetContact.ContactStatus),
			IsFavorite:      targetContact.IsFavorite,
			TagIds:          targetContact.TagIds,
		}
	}
	return &service.GetContactsResponse{Contacts: contacts}, nil
}

func (t TalkHandler) UpdateContact(ctx context.Context, request *service.UpdateContactRequest) (*service.UpdateContactResponse, error) {
	mid := utils.GetMid(ctx)
	_, err := dbClient.FetchUser(request.Contact.Mid)
	if err != nil {
		return &service.UpdateContactResponse{}, err
	}
	err = dbClient.InsertUserContact(mid, &talkDatabase.Contact{
		MID:             request.Contact.Mid,
		OverWrittenName: request.Contact.OverWrittenName,
		TagIds:          request.Contact.TagIds,
		ContactStatus:   talkDatabase.DBContactStatusFromRPC(request.Contact.ContactStatus),
	})
	if err != nil {
		return &service.UpdateContactResponse{}, err
	}

	return &service.UpdateContactResponse{}, nil
}
