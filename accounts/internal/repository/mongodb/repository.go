package repository

import (
	"accounts/internal/model"
	"accounts/internal/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AccountRepository struct {
	accountCollection *mongo.Collection
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
		accountCollection: db.Collection(collection, &options.CollectionOptions{
			Registry: reg,
		}),
	}
}

func (ar *AccountRepository) Create(ctx context.Context, account model.Account) (string, error) {
	result, err := ar.accountCollection.InsertOne(ctx, account)
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

	result, err := ar.accountCollection.UpdateOne(ctx, filter, update)
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
	cur, err := ar.accountCollection.Find(ctx, primitive.D{{}})
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

	result := ar.accountCollection.FindOne(ctx, filter)
	if result.Err() != nil {
		utils.Logger.Info("FindOne from GetOne error")
		return acc, fmt.Errorf("faild to find one account by id %s", account.ID)
	}

	if err = result.Decode(&acc); err != nil {
		return acc, fmt.Errorf("faild to decode account by id %s", account.ID)
	}
	return acc, nil
}

func (ar *AccountRepository) UpdateWithTrx(ctx context.Context, accountID, amountOfChange, operationType string) (model.Account, error) {
	var updatedAccount = &model.Account{}

	err := ar.accountCollection.Database().Client().UseSession(ctx, func(sc mongo.SessionContext) error {
		err := sc.StartTransaction(options.Transaction().
			SetReadConcern(readconcern.Snapshot()).
			SetWriteConcern(writeconcern.New(writeconcern.WMajority())),
		)

		if err != nil {
			return err
		}

		objectId, err := primitive.ObjectIDFromHex(accountID)
		if err != nil {
			return err
		}

		//select
		filter := bson.M{"_id": objectId}

		var acc model.Account
		err = ar.accountCollection.FindOne(sc, filter).Decode(&acc)
		if err == mongo.ErrNoDocuments {
			if err := sc.AbortTransaction(sc); err != nil {
				utils.Logger.Info("aborting transaction error: ", err)
			}
			utils.Logger.Info("caught exception during transaction, aborting.", err)
			return fmt.Errorf("faild to decode account by id %s", accountID)
		}

		newBalance, err := ChangeBalance(acc.Balance, amountOfChange, operationType)
		if err != nil {
			if err := sc.AbortTransaction(sc); err != nil {
				utils.Logger.Info("aborting transaction error: ", err)
			}
			utils.Logger.Info("caught exception during transaction, aborting.", err)
			return fmt.Errorf("faild to decode account by id %s", accountID)
		}

		updatedAccount = &model.Account{
			ID:       acc.ID,
			Email:    acc.Email,
			WalletID: acc.WalletID,
			Balance:  newBalance,
		}

		accoutnBytes, err := bson.Marshal(updatedAccount)
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

		//update
		update := bson.M{
			"$set": updateAccountObject,
		}

		result, _ := ar.accountCollection.UpdateOne(sc, filter, update)
		if err != nil {
			if err := sc.AbortTransaction(sc); err != nil {
				utils.Logger.Info("aborting transaction error: ", err)
			}
			utils.Logger.Info("caught exception during transaction, aborting.", err)
			return err
		}

		if result.MatchedCount == 0 {
			return fmt.Errorf("not found")
		}

		// конец

		for {
			err = sc.CommitTransaction(sc)
			switch e := err.(type) {
			case nil:
				return nil
			case mongo.CommandError:
				if e.HasErrorLabel("UnknownTransactionCommitResult") {
					log.Println("UnknownTransactionCommitResult, retrying commit operation...")
					continue
				}
				log.Println("Error during commit...")
				return e
			default:
				log.Println("Error during commit...")
				return e
			}
		}
	})
	if err != nil {
		return *updatedAccount, err
	}
	return *updatedAccount, nil
}

func ChangeBalance(oldBalance, amountOfChange string, operation string) (string, error) {
	obToFloat, _ := strconv.ParseFloat(oldBalance, 32)
	chToFloat, _ := strconv.ParseFloat(amountOfChange, 32)

	if operation == "withdrawal" && obToFloat < chToFloat {
		return oldBalance, fmt.Errorf("not enough balance")
	}

	var res float64
	switch operation {
	case "withdrawal":
		res = obToFloat - chToFloat

	case "deposit":
		res = obToFloat + chToFloat

	}

	return fmt.Sprint(res), nil
}
