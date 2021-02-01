package talkDatabase

import (
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl *DBClient) InsertNewUser(user *User) error {
	_, err := cl.UserCol.InsertOne(cl.Ctx, user)
	return err
}

func (cl *DBClient) FetchUser(mid string) (*User, error) {
	rs := cl.UserCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", mid}},
	)
	var user *User
	if rs.Decode(&user) != nil {
		return user, status.New(codes.NotFound, "user not found").Err()
	}
	return user, nil
}

func (cl *DBClient) FetchUserAttribute(mid string, attributes bson.D) (*User, error) {
	rs := cl.UserCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", mid}},
		options.FindOne().SetProjection(attributes),
	)
	var user *User
	if rs.Decode(&user) != nil {
		return user, status.New(codes.NotFound, "user not found").Err()
	}
	return user, nil
}

func (cl *DBClient) FetchUserSetting(mid string) (*Setting, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"setting", 1}})
	if rs == nil {
		return nil, err
	}
	return &rs.Setting, err
}

func (cl *DBClient) FetchUserProfile(mid string) (*Profile, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"setting", 1}})
	if rs == nil {
		return nil, err
	}
	return &rs.Profile, err
}

func (cl *DBClient) FetchUserContact(mid, contactMid string) (*Contact, bool) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"contacts", 1}})
	if err != nil {
		return nil, false
	}
	value, ok := rs.Contacts[contactMid]
	if !ok {
		value = Contact{MID: contactMid}
	}
	return &value, ok
}

func (cl *DBClient) FetchUserTag(mid, tagId string) (*Tag, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"tags", 1}})
	if err != nil {
		return nil, err
	}
	value, ok := rs.Tags[tagId]
	if !ok {
		return nil, status.New(codes.NotFound, "tag not found").Err()
	}
	return &value, nil
}

func (cl *DBClient) FetchUserGroupSettings(mid, gid string) (*GroupSetting, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"groupSettings", 1}})
	if err != nil {
		return nil, err
	}
	value, ok := rs.GroupSettings[gid]
	if !ok {
		return nil, err
	}
	return &value, nil
}

func (cl *DBClient) FetchUserJoinedGroupIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"jGroupIds", 1}})
	if err != nil {
		return []string{}, err
	}
	return rs.JoinedGroupIds, err
}

func (cl *DBClient) FetchUserInvitedGroupIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"iGroupIds", 1}})
	if err != nil {
		return []string{}, err
	}
	return rs.InvitedGroupIds, err
}

func (cl *DBClient) FetchUserFriendIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"friendIds", 1}})
	if err != nil {
		return []string{}, err
	}
	return rs.FriendIds, err
}

func (cl *DBClient) FetchUserBlockedIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"blockedIds", 1}})
	if err != nil {
		return []string{}, err
	}
	return rs.BlockedIds, err
}

func (cl *DBClient) FetchUserHashedPassword(mid string) (string, error) {
	rs, err := cl.FetchUserAttribute(mid, bson.D{{"hashedPasswd", 1}})
	if err != nil {
		return "", err
	}
	return rs.HashedPassword, err
}

func (cl *DBClient) UpdateUser(mid string, attrToUpdate bson.D) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$set": attrToUpdate})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

func (cl *DBClient) DeleteUserTag(mid, tagID string) error {
	_, err := cl.UserCol.UpdateOne(
		cl.Ctx, bson.M{"_id": mid},
		bson.M{"$unset": "tags." + tagID})
	if err != nil {
		return status.New(codes.NotFound, "tag not found").Err()
	}
	return nil
}

func (cl *DBClient) InsertOrUpdateUserTag(mid string, tag Tag) error {
	err := cl.UpdateUser(mid, bson.D{{
		"tags." + tag.TagID, tag,
	}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

func (cl *DBClient) InsertOrUpdateUserContact(mid string, contact *Contact) error {
	err := cl.UpdateUser(mid, bson.D{{
		"contacts." + contact.MID, contact,
	}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

func (cl *DBClient) UpdateUserContactStatus(mid, targetMid string, status service.ContactStatus) error {
	_, ok := cl.FetchUserContact(mid, targetMid)
	if ok {
		err := cl.UpdateUser(mid, bson.D{{"contacts." + targetMid + ".cStatus", status}})
		return err
	}
	err := cl.InsertOrUpdateUserContact(mid, &Contact{
		MID:           targetMid,
		TagIds:        []string{},
		ContactStatus: RPCContactStatusToDBContactStatus(status),
	})
	return err
}

func (cl *DBClient) AddToSetUserAttribute(mid, fieldName string, object interface{}) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$addToSet": bson.M{fieldName: bson.M{"$each": object}}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

func (cl *DBClient) RemoveFromSetUserAttribute(mid, fieldName, target string) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$pull": bson.M{fieldName: target}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

func (cl *DBClient) RemoveFromSetUserAttributes(mid string, targets bson.M) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$pull": targets})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}
