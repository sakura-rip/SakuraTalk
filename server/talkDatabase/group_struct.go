package talkDatabase

type Group struct {
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
