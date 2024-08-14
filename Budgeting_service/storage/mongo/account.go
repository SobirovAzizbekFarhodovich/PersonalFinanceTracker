package mongo

import (
	"context"
	"errors"
	"fmt"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Account struct {
	mongo *mongo.Collection
}

func NewAccount(db *mongo.Database) *Account {
	return &Account{mongo: db.Collection("account")}
}

func (s *Account) CreateAccount(req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	bsonAccount := e.AccountToBSON(req.Account)
	_, err := s.mongo.InsertOne(context.TODO(), bsonAccount)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{}, nil
}

func (s *Account) UpdateAccount(req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	bsonAccount := e.AccountToBSON(req.Account)
	filter := bson.M{"_id": bsonAccount.ID}
	update := bson.M{"$set": bsonAccount}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateAccountResponse{}, nil
}

func (s *Account) DeleteAccount(req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	_, err := s.mongo.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountResponse{}, nil
}

func (s *Account) GetAccount(req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	var bsonAccount m.Account
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonAccount)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("Account not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: e.BsonToAccount(&bsonAccount)}, nil
}

func (s *Account) ListAccounts(req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Account
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Account
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToAccount(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListAccountsResponse{
		Accounts: accounts,
	}, nil
}

func (s *Account) GetAmount(req *pb.GetAmountRequest) (*pb.GetAmountResponse, error) {
	filter := bson.M{"user_id": req.UserId}

	var account m.Account

	err := s.mongo.FindOne(context.TODO(), filter).Decode(&account)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("account not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetAmountResponse{
		Balance: account.Balance,
	}, nil
}

func (s *Account) UpdateAmount(req *pb.UpdateAmountRequest) (*pb.UpdateAmountResponse, error) {
	filter := bson.M{"user_id": req.UserId}
	update := bson.M{
		"$set": bson.M{
			"balance": req.Balance,
		},
	}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}

	return &pb.UpdateAmountResponse{}, nil
}
