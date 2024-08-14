package mongo

import (
	"context"
	"errors"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

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
	var bsonAccount m.Transaction
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonAccount)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("Transaction not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetTransactionResponse{Transaction: e.BsonToTransaction(&bsonAccount)}, nil
}

func (s *Transaction) ListTransactions(req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Transaction
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Transaction
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToTransaction(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListTransactionsResponse{
		Transactions: accounts,
	}, nil
}

func (s *Transaction) Spending(req *pb.SpendingRequest) (*pb.SpendingResponse, error) {
	filter := bson.M{"user_id": req.UserId, "type": "expense"}

	cursor, err := s.mongo.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var spendingCount int32
	var spendingMoney float32

	for cursor.Next(context.TODO()) {
		var bsonTransaction m.Transaction
		if err := cursor.Decode(&bsonTransaction); err != nil {
			return nil, err
		}

		spendingCount++
		spendingMoney += bsonTransaction.Amount
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.SpendingResponse{
		SpendingCount: spendingCount,
		SpendingMoney: spendingMoney,
	}, nil
}

func (s *Transaction) Income(req *pb.IncomeRequest) (*pb.IncomeResponse, error) {
	filter := bson.M{"user_id": req.UserId, "type": "income"}

	cursor, err := s.mongo.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var incomeCount int32
	var incomeMoney float32

	for cursor.Next(context.TODO()) {
		var bsonTransaction m.Transaction
		if err := cursor.Decode(&bsonTransaction); err != nil {
			return nil, err
		}

		incomeCount++
		incomeMoney += bsonTransaction.Amount
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.IncomeResponse{
		IncomeCount: incomeCount,
		IncomeMoney: incomeMoney,
	}, nil
}
