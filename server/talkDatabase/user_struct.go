package talkDatabase

type User struct {
	MID            string `bson:"_id"`
	HashedPassword string `bson:"hashedPasswd"`

	Profile Profile `bson:"profile"`
	Setting Setting `bson:"setting"`

	JoinedGroupIds  []string `bson:"jGroupIds"`
	InvitedGroupIds []string `bson:"iGroupIds"`
	FriendIds       []string `bson:"friendIds"`
	BlockedIds      []string `bson:"blockedIds"`

	GroupSettings map[string]GroupSetting `bson:"groupSettings"`
	Tags          map[string]Tag          `bson:"tags"`
	Contacts      map[string]Contact      `bson:"contacts"`
}

func NewUser() *User {
	return &User{}
}

type GroupSetting struct {
	//新規メッセージの通知など
	EnableNotification bool `bson:"enNotify"`
	//メンションの通知
	EnableNotificationMention bool `bson:"enNotifyMention"`
	//誰かが新規参加した時の通知
	EnableNotificationOnJoin bool `bson:"enNotifyJoin"`
	//誰かが誰かをキックした時の通知
	EnableNotificationOnKick bool `bson:"enNotifyKick"`

	OverWrittenName string   `bson:"owName"`
	TagIds          []string `bson:"tagIds"`
}

func NewGroupSetting() *GroupSetting {
	return &GroupSetting{}
}

type Profile struct {
	Name      string `bson:"name"`
	Bio       string `bson:"bio"`
	IconPath  string `bson:"iPath"`
	CoverPath string `bson:"cPath"`

	Location string `bson:"location"`
	WebSite  string `bson:"website"`
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
	AllowSearchByEmail         bool `bson:"asByUserID"`
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

	ContactStatus int64 `bson:"cStatus"`
}

func NewContact() *Contact {
	return &Contact{}
}
