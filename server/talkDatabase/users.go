package talkDatabase

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl *DBClient) InsertNewUser(user *User) error {
	_, err := cl.UserCol.InsertOne(cl.Ctx, user)
	return err
}

func (cl *DBClient) InsertNewGroup(group *Group) error {
	_, err := cl.GroupCol.InsertOne(cl.Ctx, group)
	return err
}

func (cl *DBClient) FetchUserAttribute(mid string, attributes bson.D) (*User, error) {
	rs := cl.UserCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", mid}},
		options.FindOne().SetProjection(attributes),
	)
	var user *User
	if rs.Decode(&user) != nil {
		return user, status.New(codes.NotFound, "use not found").Err()
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

func (cl *DBClient) FetchUserContact(mid, contactMid string) (*Contact, error) {
	otherProfile, err := cl.FetchUserProfile(contactMid)
	if err != nil {
		return nil, err
	}
	contact := &Contact{
		MID:           contactMid,
		ContactStatus: 0,
	}
	user, err := cl.FetchUserAttribute(mid, bson.D{{"contacts", 1}})
	if err != nil {
		return nil, err
	}
	if len(user.Contacts) == 0 {
		contact.ContactStatus = 1
	}
}
