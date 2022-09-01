package mongodb

import (
	"accounts/internal/db/model"
	"accounts/internal/db/storage"
	"accounts/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type db struct {
	collection *mongo.Collection
	logger     *zap.SugaredLogger
}

func (d *db) Create(ctx context.Context, account model.Account) (string, error) {
	result, err := d.collection.InsertOne(ctx, account)
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", err
}

func (d *db) Update(ctx context.Context, account model.Account) error {
	objectId, err := primitive.ObjectIDFromHex(account.ID)
	if err != nil {
		return err
	}

	//select
	filter := bson.M{"_id": objectId}
	//filter := bson.M{"email": account.Email}

	accoutnBytes, err := bson.Marshal(account)
	if err != nil {
		utils.Logger.Info("bson marhal error ", err)
		return err
	}

	var updateAccountObject bson.M
	err = bson.Unmarshal(accoutnBytes, &updateAccountObject)
	if err != nil {
		utils.Logger.Info("bson unmarhal error ", err)
		return err
	}

	delete(updateAccountObject, "_id")
	//delete(updateAccountObject, "email")

	//update
	update := bson.M{
		"$set": updateAccountObject,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.Logger.Info("update one error ", err)
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}
	// TODO return New Balance
	return nil
}

func (d *db) GetOne(ctx context.Context, account model.Account) (a model.Account, err error) {
	objectId, err := primitive.ObjectIDFromHex(account.ID)
	if err != nil {
		utils.Logger.Info("error convert account id")
	}

	//select
	filter := bson.M{"_id": objectId}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return a, fmt.Errorf("faild to find one account by id %s", account.ID)
	}

	if err = result.Decode(&a); err != nil {
		return a, fmt.Errorf("faild to decode account by id %s", account.ID)
	}
	return a, nil
}

func NewStorage(database *mongo.Database, collection string, logger *zap.SugaredLogger) storage.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
