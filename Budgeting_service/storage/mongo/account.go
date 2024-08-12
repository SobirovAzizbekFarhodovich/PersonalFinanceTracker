package mongo

import (
	"context"
	"errors"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		return nil, errors.New("progress not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetAccountResponse{Account: e.BsonToAccount(&bsonAccount)}, nil
}

func (s *Account) ListAccounts(req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	pipeline := []bson.M{
		{"$skip": skip},
		{"$limit": limit},
		{"$project": bson.M{
			"_id":      0,
			"user_id":  1,
			"name":     1,
			"type":     1,
			"balance":  1,
			"currency": 1,
		}},
	}

	cursor, err := s.mongo.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Account
	for cursor.Next(context.TODO()) {
		var account pb.Account
		if err := cursor.Decode(&account); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return &pb.ListAccountsResponse{
		Accounts: accounts,
	}, nil
}
