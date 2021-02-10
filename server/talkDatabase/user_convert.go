package talkDatabase

import service "github.com/sakura-rip/SakuraTalk/talkService"

func (cl *Profile) ConvertToRPCStruct() *service.Profile {
	rpcProfile := &service.Profile{
		DisplayName: cl.Name,
		Bio:         cl.Bio,
		IconPath:    cl.IconPath,
		CoverPath:   cl.CoverPath,
		TwitterID:   cl.TwitterID,
		InstagramID: cl.InstagramID,
		GithubID:    cl.GithubID,
		WebSiteURL:  cl.WebSiteURL,
		Location:    cl.Location,
		BirthDay:    cl.BirthDay,
		Status:      cl.Status,
	}
	return rpcProfile
}

func (cl *Setting) ConvertToRPCStruct() *service.Setting {
	rpcSetting := &service.Setting{
		PrivateUserID:              cl.PrivateUserID,
		Email:                      cl.Email,
		UserTicket:                 cl.UserTicket,
		AllowSearchByPrivateUserID: cl.AllowSearchByPrivateUserID,
		AllowSearchByEmail:         cl.AllowSearchByEmail,
		AllowSearchByUserTicket:    cl.AllowSearchByUserTicket,
	}
	return rpcSetting
}

func (cl *Tag) ConvertToRPCStruct() *service.Tag {
	rpcTag := &service.Tag{
		TagID:       cl.TagID,
		Name:        cl.Name,
		Description: cl.Description,
		Color:       cl.Color,
		Creator:     cl.Creator,
		CreatedTime: cl.CreatedTime,
	}
	return rpcTag
}

func DBContactStatusFromRPC(rpcCs service.ContactStatus) int64 {
	return int64(rpcCs)
}

func RPCContactStatusFromDB(dbCs int64) service.ContactStatus {
	return service.ContactStatus(dbCs)
}

func (cl *Contact) ConvertToRPCStruct() *service.Contact {
	rpcContact := &service.Contact{
		Mid:             cl.MID,
		OverWrittenName: cl.OverWrittenName,
		ContactStatus:   RPCContactStatusFromDB(cl.ContactStatus),
		IsFavorite:      cl.IsFavorite,
		TagIds:          cl.TagIds,
	}
	return rpcContact
}
