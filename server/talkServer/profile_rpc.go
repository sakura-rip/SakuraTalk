package talkServer

import (
	"context"
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"github.com/sakura-rip/SakuraTalk/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func (t TalkHandler) UpdateProfileAttributes(ctx context.Context, request *service.UpdateProfileAttributesRequest) (*service.UpdateProfileAttributesResponse, error) {
	var updateObject bson.D
	for _, key := range request.Keys {
		switch key {
		case service.ProfileKey_DISPLAY_NAME:
			updateObject = append(updateObject, bson.E{Key: "profile.name", Value: request.Profile.DisplayName})
		case service.ProfileKey_BIO:
			updateObject = append(updateObject, bson.E{Key: "profile.bio", Value: request.Profile.Bio})
		case service.ProfileKey_ICON_PATH:
			updateObject = append(updateObject, bson.E{Key: "profile.iPath", Value: request.Profile.IconPath})
		case service.ProfileKey_COVER_PATH:
			updateObject = append(updateObject, bson.E{Key: "profile.cPath", Value: request.Profile.CoverPath})
		case service.ProfileKey_TWITTER_ID:
			updateObject = append(updateObject, bson.E{Key: "profile.twitterID", Value: request.Profile.TwitterID})
		case service.ProfileKey_INSTAGRAM_ID:
			updateObject = append(updateObject, bson.E{Key: "profile.instagramID", Value: request.Profile.InstagramID})
		case service.ProfileKey_GITHUB_ID:
			updateObject = append(updateObject, bson.E{Key: "profile.githubID", Value: request.Profile.GithubID})
		case service.ProfileKey_WEB_SITE_URL:
			updateObject = append(updateObject, bson.E{Key: "profile.webSiteURL", Value: request.Profile.WebSiteURL})
		case service.ProfileKey_LOCATION:
			updateObject = append(updateObject, bson.E{Key: "profile.location", Value: request.Profile.Location})
		case service.ProfileKey_BIRTH_DAY:
			updateObject = append(updateObject, bson.E{Key: "profile.birthday", Value: request.Profile.BirthDay})
		case service.ProfileKey_STATUS:
			updateObject = append(updateObject, bson.E{Key: "profile.status", Value: request.Profile.Status})
		}
	}
	err := dbClient.UpdateUser(utils.GetMid(ctx), updateObject)
	if err != nil {
		return nil, err
	}
	return &service.UpdateProfileAttributesResponse{}, nil
}

func (t TalkHandler) GetProfile(ctx context.Context, empty *service.Empty) (*service.Profile, error) {
	profile, err := dbClient.FetchUserProfile(utils.GetMid(ctx))
	if err != nil {
		return nil, err
	}
	return profile.ConvertToRPCStruct(), nil
}
