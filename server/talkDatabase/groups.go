package talkDatabase

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (cl *DBClient) InsertNewGroup(group *Group) error {
	_, err := cl.GroupCol.InsertOne(cl.Ctx, group)
	return err
}

func (cl *DBClient) FetchGroup(gid string) (*Group, error) {
	rs := cl.GroupCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", gid}},
	)
	var grp *Group
	if rs.Decode(&grp) != nil {
		return grp, status.New(codes.NotFound, "use not found").Err()
	}
	return grp, nil
}

func (cl *DBClient) FetchGroupAttributes(gid string, attributes bson.D) (*Group, error) {
	rs := cl.GroupCol.FindOne(
		cl.Ctx,
		bson.D{{"_id", gid}},
		options.FindOne().SetProjection(attributes),
	)
	var grp *Group
	if rs.Decode(&grp) != nil {
		return grp, status.New(codes.NotFound, "use not found").Err()
	}
	return grp, nil
}
