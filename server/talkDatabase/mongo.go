package talkDatabase

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBClient struct {
	Session  *mongo.Client
	UserCol  *mongo.Collection
	GroupCol *mongo.Collection
}

func NewDBClient() *DBClient {
	cl := &DBClient{
		Session: ConnectToMongoDB(),
	}
	cl.UserCol = cl.Session.Database("sakuraTalk").Collection("users")
	cl.GroupCol = cl.Session.Database("sakuraTalk").Collection("groups")
	return cl
}

// ConnectToMongoDB connect to mongo
func ConnectToMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongo.Connect(ctx, options.Client().ApplyURI("mangodb://localhost:27017"))
	if err != nil {
		panic(fmt.Errorf("error on init mongo db: %v", err))
	}
	return db
}
