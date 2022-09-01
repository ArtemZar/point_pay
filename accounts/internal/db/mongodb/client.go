package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port, username, password, database, authDB string) (db *mongo.Database, err error) {
	var connectionString string
	var isAuth bool

	if username == "" && password == "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clitentOtions := options.Client().ApplyURI(connectionString)
	if isAuth {
		if authDB == "" {
			authDB = database
		}
		clitentOtions.SetAuth(options.Credential{
			AuthSource: authDB,
			Username:   username,
			Password:   password,
		})
	}

	client, err := mongo.Connect(ctx, clitentOtions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	return client.Database(database), nil
}
