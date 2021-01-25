package talkDatabase

func (cl *DBClient) InsertNewGroup(group *Group) error {
	_, err := cl.GroupCol.InsertOne(cl.Ctx, group)
	return err
}
