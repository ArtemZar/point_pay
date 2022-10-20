package repository

import (
	"accounts/internal/model"
	"accounts/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type AccountRepository struct {
	db *mongo.Collection
}

func NewAccountRepository(db *mongo.Database, collection string) *AccountRepository {

	// create a custom registry builder
	rb := bsoncodec.NewRegistryBuilder()

	// register default codecs and encoders/decoders
	var primitiveCodecs bson.PrimitiveCodecs
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(rb)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(rb)
	primitiveCodecs.RegisterPrimitiveCodecs(rb)

	// register custom encoder/decoder
	myNumberType := reflect.TypeOf(model.MyNumber(0))

	rb.RegisterTypeEncoder(
		myNumberType,
		bsoncodec.ValueEncoderFunc(func(_ bsoncodec.EncodeContext, vw bsonrw.ValueWriter, val reflect.Value) error {
			if !val.IsValid() || val.Type() != myNumberType {
				return bsoncodec.ValueEncoderError{
					Name:     "MyNumberEncodeValue",
					Types:    []reflect.Type{myNumberType},
					Received: val,
				}
			}
			// IMPORTANT STEP: cast uint64 to int64 so it can be stored in mongo
			vw.WriteInt64(int64(val.Uint()))
			return nil
		}),
	)

	rb.RegisterTypeDecoder(
		myNumberType,
		bsoncodec.ValueDecoderFunc(func(_ bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
			// IMPORTANT STEP: read sore value in mongo as int64
			read, err := vr.ReadInt64()
			if err != nil {
				return err
			}
			// IMPORTANT STEP: cast back to uint64
			val.SetUint(uint64(read))
			return nil
		}),
	)

	// build the registry
	reg := rb.Build()

	return &AccountRepository{
		db: db.Collection(collection, &options.CollectionOptions{
			Registry: reg,
		}),
	}
}

func (ar *AccountRepository) Create(ctx context.Context, account model.Account) (string, error) {
	result, err := ar.db.InsertOne(ctx, account)
	if err != nil {
		return "", err
	}
	objectId, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return objectId.Hex(), nil
	}
	return "", err
}

func (ar *AccountRepository) Update(ctx context.Context, account model.Account) error {
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

	result, err := ar.db.UpdateOne(ctx, filter, update)
	if err != nil {
		utils.Logger.Info("update one error ", err)
		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

func (ar *AccountRepository) Find(ctx context.Context) (acc []model.Account, err error) {
	cur, err := ar.db.Find(ctx, primitive.D{{}})
	if err != nil {
		utils.Logger.Error(err)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		data := &model.Account{}
		err := cur.Decode(data)
		if err != nil {
			utils.Logger.Error(err)
		}
		acc = append(acc, *data)

	}
	if err := cur.Err(); err != nil {
		utils.Logger.Error(err)
	}
	return acc, nil
}

func (ar *AccountRepository) GetOne(ctx context.Context, account model.Account) (acc model.Account, err error) {
	objectId, err := primitive.ObjectIDFromHex(account.ID)
	if err != nil {
		utils.Logger.Info("error convert account id")
	}

	//select
	filter := bson.M{"_id": objectId}

	result := ar.db.FindOne(ctx, filter)
	if result.Err() != nil {
		return acc, fmt.Errorf("faild to find one account by id %s", account.ID)
	}

	if err = result.Decode(&acc); err != nil {
		return acc, fmt.Errorf("faild to decode account by id %s", account.ID)
	}
	return acc, nil
}
