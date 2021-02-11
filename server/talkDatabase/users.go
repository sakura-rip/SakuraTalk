package talkDatabase

import (
	service "github.com/sakura-rip/SakuraTalk/talkService"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//InsertNewUser 新しいUserをデータベースに保存します
func (cl *DBClient) InsertNewUser(user *User) error {
	_, err := cl.UserCol.InsertOne(cl.Ctx, user)
	return err
}

//FetchUser MIDから、Userのデータをデータベースから取り出します
func (cl *DBClient) FetchUser(mid string) (*User, error) {
	rs := cl.UserCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", mid}},
	)
	var user *User
	if rs.Decode(&user) != nil {
		return user, status.Error(codes.NotFound, "user not found")
	}
	return user, nil
}

//FetchUserAttribute Userの中の特定の情報：attributesを取り出します
func (cl *DBClient) FetchUserAttribute(mid string, attributes bson.D) (*User, error) {
	rs := cl.UserCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", mid}},
		options.FindOne().SetProjection(attributes),
	)
	var user *User
	if rs.Decode(&user) != nil {
		return user, status.Error(codes.NotFound, "user not found")
	}
	return user, nil
}

func (cl *DBClient) FetchUserAttributes(mid string, attributes ...string) (*User, error) {
	var projections = bson.D{}
	for _, attribute := range attributes {
		projections = append(projections, bson.E{Key: attribute, Value: 1})
	}
	return cl.FetchUserAttribute(mid, projections)
}

//FetchUserSetting　データベースの中からMIDのユーザーのSettingを取り出します
func (cl *DBClient) FetchUserSetting(mid string) (*Setting, error) {
	rs, err := cl.FetchUserAttributes(mid, "setting")
	if rs == nil {
		return nil, err
	}
	return &rs.Setting, err
}

//FetchUserProfile データーベースの中からMIDのユーザーのプロフォールを取得する
func (cl *DBClient) FetchUserProfile(mid string) (*Profile, error) {
	rs, err := cl.FetchUserAttributes(mid, "profile")
	if rs == nil {
		return nil, err
	}
	return &rs.Profile, err
}

//FetchUserContact　データーベースのなかからMIDごとのContactを取得する
//存在しない場合、特にデフォルト値をいじらなければいけない部分はないので、そのままMIDをつけたContactを返す
func (cl *DBClient) FetchUserContact(mid, contactMid string) (*Contact, error) {
	rs, err := cl.FetchUserAttributes(mid, "contacts")
	if err != nil {
		return nil, err
	}
	value, ok := rs.Contacts[contactMid]
	if !ok {
		value = Contact{MID: contactMid}
	}
	return &value, nil
}

//FetchUserTag　データベースのなかからMIDごとのTAGを取得する
//存在しない場合はerrorを返す
func (cl *DBClient) FetchUserTag(mid, tagId string) (*Tag, error) {
	rs, err := cl.FetchUserAttributes(mid, "tags")
	if err != nil {
		return nil, err
	}
	value, ok := rs.Tags[tagId]
	if !ok {
		return nil, status.Error(codes.NotFound, "tag not found")
	}
	return &value, nil
}

//FetchUserGroupSettings　ユーザーごとのGroupの設定を取得する。
//存在しない場合は、デフォルトの設定を返す
func (cl *DBClient) FetchUserGroupSettings(mid, gid string) (*GroupSetting, bool) {
	rs, err := cl.FetchUserAttributes(mid, "groupSettings."+gid)
	if err != nil {
		return &GroupSetting{
			EnableNotification:        false,
			EnableNotificationMention: true,
			EnableNotificationOnJoin:  true,
			EnableNotificationOnKick:  true,
			OverWrittenName:           "",
			TagIds:                    []string{},
			IsFavorite:                false,
		}, false
	}
	value, _ := rs.GroupSettings[gid]
	return &value, true
}

//FetchUserJoinedGroupIds　データベースから、MIDの参加しているGroupのMID一覧を取得する
func (cl *DBClient) FetchUserJoinedGroupIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttributes(mid, "jGroupIds")
	if err != nil {
		return []string{}, err
	}
	return rs.JoinedGroupIds, err
}

//FetchUserInvitedGroupIds　デーたーべすから、MIDの招待されているGroupのMID一覧を取得する
func (cl *DBClient) FetchUserInvitedGroupIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttributes(mid, "iGroupIds")
	if err != nil {
		return []string{}, err
	}
	return rs.InvitedGroupIds, err
}

//FetchUserFriendIds データベースから、MIDの友達のID一覧を取得する
func (cl *DBClient) FetchUserFriendIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttributes(mid, "friendIds")
	if err != nil {
		return []string{}, err
	}
	return rs.FriendIds, err
}

//FetchUserBlockedIds　データベースから、MIDがブロックしている人の一覧を取得する
func (cl *DBClient) FetchUserBlockedIds(mid string) ([]string, error) {
	rs, err := cl.FetchUserAttributes(mid, "blockedIds")
	if err != nil {
		return []string{}, err
	}
	return rs.BlockedIds, err
}

//FetchUserHashedPassword　データベースから、ハッシュ化されているパスワードを取り出す。
func (cl *DBClient) FetchUserHashedPassword(mid string) (string, error) {
	rs, err := cl.FetchUserAttributes(mid, "hashedPasswd")
	if err != nil {
		return "", err
	}
	return rs.HashedPassword, err
}

//UpdateUser　MIDのattrToUpdateを更新する
func (cl *DBClient) UpdateUser(mid string, attrToUpdate bson.D) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$set": attrToUpdate})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//DeleteUserTag　ユーザーのTAGをtagIDから削除する
func (cl *DBClient) DeleteUserTag(mid, tagID string) error {
	_, err := cl.UserCol.UpdateOne(
		cl.Ctx, bson.M{"_id": mid},
		bson.M{"$unset": "tags." + tagID})
	if err != nil {
		return status.Error(codes.NotFound, "tag not found")
	}
	return nil
}

//InsertUserTag TAGをデータベースに保存します。
//すでにある場合は上書きします
func (cl *DBClient) InsertUserTag(mid string, tag Tag) error {
	err := cl.UpdateUser(mid, bson.D{{
		"tags." + tag.TagID, tag,
	}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//InsertUserContact　MIDごとのContactを、データベースに保存します。
//すでにある場合は上書きします
func (cl *DBClient) InsertUserContact(mid string, contact *Contact) error {
	err := cl.UpdateUser(mid, bson.D{{
		"contacts." + contact.MID, contact,
	}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//UpdateUserContactStatus　MIDごとのContactのContactStatusを更新します。
//データベースにContactが存在しない場合、デフォルト値で新たに作ります
func (cl *DBClient) UpdateUserContactStatus(mid, targetMid string, status service.ContactStatus) error {
	_, err := cl.FetchUserContact(mid, targetMid)
	if err != nil {
		err := cl.UpdateUser(mid, bson.D{{"contacts." + targetMid + ".cStatus", status}})
		return err
	}
	err = cl.InsertUserContact(mid, &Contact{
		MID:           targetMid,
		TagIds:        []string{},
		ContactStatus: DBContactStatusFromRPC(status),
	})
	return err
}

//UpdateUserContactIsFavorite MIDのユーザーのContactの中の、targetMIDのisFavoriteを変更する。
//すでに追加/ブロック/削除されていることが前提なのでそのままＵｐｄａｔｅする
func (cl *DBClient) UpdateUserContactIsFavorite(mid, targetMid string, tOrF bool) error {
	err := cl.UpdateUser(mid, bson.D{{"contacts." + targetMid + ".isFavorite", tOrF}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//AddToSetUserAttribute UserのArray要素に、データを追加します。
//すでに存在している場合は、追加しません
func (cl *DBClient) AddToSetUserAttribute(mid, fieldName string, object interface{}) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$addToSet": bson.M{fieldName: object}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//AddToSetUserAttributes　Userの複数のArray要素にデータを追加します
//すでに存在している場合は追加しません
func (cl *DBClient) AddToSetUserAttributes(mid, object bson.M) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$addToSet": object})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//RemoveFromSetUserAttribute UserのArray要素から、データを削除します。
func (cl *DBClient) RemoveFromSetUserAttribute(mid, fieldName, target string) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$pull": bson.M{fieldName: target}})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}

//RemoveFromSetUserAttribute Userの複数のArray要素から、データを削除します。
func (cl *DBClient) RemoveFromSetUserAttributes(mid string, targets bson.M) error {
	_, err := cl.UserCol.UpdateOne(cl.Ctx, bson.M{"_id": mid}, bson.M{"$pull": targets})
	if err != nil {
		return status.Error(codes.Internal, "db error")
	}
	return nil
}
