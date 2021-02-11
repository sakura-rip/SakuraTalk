package talkDatabase

type Group struct {
	GID         string   `bson:"_id"`
	CreatorMID  string   `bson:"creator"`
	CreatedTime int64    `bson:"createdTime"`
	MemberMids  []string `bson:"members"`
	InvitedMids []string `bson:"invited"`

	Name        string `bson:"name"`
	Description string `bson:"description"`
	IconPath    string `bson:"iPath"`
	CoverPath   string `bson:"cPath"`

	GroupTicket            string `bson:"gTicket"`
	AllowJoinByGroupTicket bool   `bson:"ajByGTicket"`
}

func NewGroup() *Group {
	return &Group{}
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
	IsFavorite      bool     `bson:"isFavorite"`
}

func NewGroupSetting() *GroupSetting {
	return &GroupSetting{
		EnableNotification:        false,
		EnableNotificationMention: true,
		EnableNotificationOnJoin:  true,
		EnableNotificationOnKick:  true,
		OverWrittenName:           "",
		TagIds:                    []string{},
		IsFavorite:                false,
	}
}
