package talkDatabase

import (
	"fmt"
	"reflect"
)

type User struct {
	MID            string `bson:"_id"`
	HashedPassword string `bson:"hashedPasswd"`

	Profile Profile `bson:"profile"`
	Setting Setting `bson:"setting"`

	JoinedGroupIds  []string `bson:"jGroupIds"`
	InvitedGroupIds []string `bson:"iGroupIds"`

	FriendIds  []string `bson:"friendIds"`
	BlockedIds []string `bson:"blockedIds"`
	DeletedIds []string `bson:"deletedIds"`

	ReceivedFriendIds  []string `bson:"recvFriendIds"`
	ReceivedBlockedIds []string `bson:"recvBlockedIds"`
	ReceivedDeletedIds []string `bson:"recvDeletedIds"`

	GroupSettings map[string]GroupSetting `bson:"groupSettings"`
	Tags          map[string]Tag          `bson:"tags"`
	Contacts      map[string]Contact      `bson:"contacts"`
}

func (cl User) GetBsonName(fieldName string) (string, error) {
	field, ok := reflect.TypeOf(cl).Elem().FieldByName(fieldName)
	if !ok {
		return "", fmt.Errorf("no such field name")
	}
	return field.Tag.Get("bson"), nil
}

func NewUser() *User {
	return &User{}
}

type Profile struct {
	Name      string `bson:"name"`
	Bio       string `bson:"bio"`
	IconPath  string `bson:"iPath"`
	CoverPath string `bson:"cPath"`

	TwitterID   string `bson:"twitterID"`
	InstagramID string `bson:"instagramID"`
	GithubID    string `bson:"githubID"`
	WebSiteURL  string `bson:"webSiteURL"`

	Location string `bson:"location"`
	BirthDay int64  `bson:"birthday"`
	Status   string `bson:"status"`
}

func NewProfile() *Profile {
	return &Profile{}
}

type Setting struct {
	PrivateUserID string `bson:"pUserID"`
	Email         string `bson:"email"`
	UserTicket    string `bson:"UTicket"`

	AllowSearchByPrivateUserID bool `bson:"asByPUserID"`
	AllowSearchByEmail         bool `bson:"asByEmail"`
	AllowSearchByUserTicket    bool `bson:"asByUserTicket"`
}

func NewSetting() *Setting {
	return &Setting{}
}

type Tag struct {
	TagID       string `bson:"tagID"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Color       string `bson:"color"`
	Creator     string `bson:"creator"`
	CreatedTime int64  `bson:"createdTime"`
}

func NewTag() *Tag {
	return &Tag{}
}

type Contact struct {
	MID             string   `bson:"mid"`
	OverWrittenName string   `bson:"owName"`
	TagIds          []string `bson:"tagIds"`

	IsFavorite    bool  `bson:"isFavorite"`
	ContactStatus int64 `bson:"cStatus"`
}

func NewContact(mid string) *Contact {
	return &Contact{MID: mid}
}
