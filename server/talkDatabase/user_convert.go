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
func RPCContactStatusToDBContactStatus(rpcCs service.ContactStatus) int64 {
	switch rpcCs {
	case service.ContactStatus_NO_RELATION:
		return 0
	case service.ContactStatus_FRIEND:
		return 1
	case service.ContactStatus_BLOCKED:
		return 2
	case service.ContactStatus_DELETED:
		return 3
	default:
		return 0
	}
}
func DBContactStatusToRPCContactStatus(dbCs int64) service.ContactStatus {
	switch dbCs {
	case 0:
		return service.ContactStatus_NO_RELATION
	case 1:
		return service.ContactStatus_FRIEND
	case 2:
		return service.ContactStatus_BLOCKED
	case 3:
		return service.ContactStatus_DELETED
	default:
		return service.ContactStatus_NO_RELATION
	}
}

func (cl *Contact) ConvertToRPCStruct() *service.Contact {
	rpcContact := &service.Contact{
		Mid:             cl.MID,
		OverWrittenName: cl.OverWrittenName,
		ContactStatus:   DBContactStatusToRPCContactStatus(cl.ContactStatus),
		TagIds:          cl.TagIds,
	}
	return rpcContact
}
