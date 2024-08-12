package mongo

import (
	"context"

	e "budgeting/extra"
	pb "budgeting/genprotos"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Transaction struct {
	mongo *mongo.Collection
}

func NewTransaction(db *mongo.Database) *Transaction {
	return &Transaction{mongo: db.Collection("Transaction")}
}

func (s *Transaction) CreateTransaction(req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	bsonTransaction := e.TransactionToBSON(req.Transaction)
	_, err := s.mongo.InsertOne(context.TODO(), bsonTransaction)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTransactionResponse{}, nil
}

func (s *Transaction) UpdateTransaction(req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	bsonTransaction := e.TransactionToBSON(req.Transaction)
	filter := bson.M{"_id": bsonTransaction.ID}
	update := bson.M{"$set": bsonTransaction}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTransactionResponse{}, nil
}

func (s *Transaction) DeleteTransaction(req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	_, err := s.mongo.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTransactionResponse{}, nil
}

func (s *Transaction) GetTransaction(req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	var budget pb.Transaction
	filter := bson.M{"id": req.Id}
	err := s.mongo.FindOne(context.TODO(), filter).Decode(&budget)
	if err != nil {
		return nil, err
	}

	var category pb.Category
	categoryFilter := bson.M{"category_id": budget.CategoryId}
	err = s.mongo.FindOne(context.TODO(), categoryFilter).Decode(&category)
	if err != nil {
		return nil, err
	}

	var account pb.Account
	accountFilter := bson.M{"account_id": budget.AccountId}
	err = s.mongo.FindOne(context.TODO(), accountFilter).Decode(&account)
	if err != nil {
		return nil, err
	}

	return &pb.GetTransactionResponse{
		Transaction: &budget,
		Category:    &category,
		Account:     &account,
	}, nil
}

func (s *Transaction) ListTransactions(req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	limit := req.Limit
	if limit == 0 {
		limit = 10
	}

	page := req.Page
	if page == 0 {
		page = 1
	}

	findOptions := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64((page - 1) * limit))

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var budgets []*pb.GetTransactionResponse
	for cursor.Next(context.TODO()) {
		var budget pb.Transaction
		if err := cursor.Decode(&budget); err != nil {
			return nil, err
		}
		var category pb.Category
		categoryFilter := bson.M{"category_id": budget.CategoryId}
		err := s.mongo.FindOne(context.TODO(), categoryFilter).Decode(&category)
		if err != nil {
			return nil, err
		}
		var account pb.Account
		accountFilter := bson.M{"account_id": budget.AccountId}
		err = s.mongo.FindOne(context.TODO(), accountFilter).Decode(&account)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, &pb.GetTransactionResponse{
			Transaction:     &budget,
			Category: &category,
			Account: &account,
		})
	}
	return &pb.ListTransactionsResponse{
		Transactions: budgets,
	}, nil
}
