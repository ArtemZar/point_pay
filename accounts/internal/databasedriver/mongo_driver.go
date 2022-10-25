package databasedriver

import (
	"accounts/internal/config"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func (mongodb *MongoDB) ConnectDatabase(cfg *config.DataBase) {
	client, err := mongo.NewClient(options.Client().ApplyURI(getConnectionString(cfg)))

	if err != nil {
		panic(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
	Mongo.Client = client
}

//func (mongodb *MongoDB) ConnectCollection(databaseName, collectionName string) *mongo.Collection {
//	return mongodb.Client.Database(databaseName).Collection(collectionName)
//}

func getConnectionString(cfg *config.DataBase) string {

	var connectionString string

	if cfg.Username == "" && cfg.Password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s", cfg.Host, cfg.Port)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port)
	}

	return connectionString
}
