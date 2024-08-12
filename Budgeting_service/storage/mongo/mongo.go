package mongo

import (
	"budgeting/storage"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB           *mongo.Database
	AccountS     storage.AccountI
	BudgetS      storage.BudgetI
	CategoryS    storage.CategoryI
	GoalS        storage.GoalI
	TransactionS storage.TransactionI
}

func ConnectMongo() (storage.StorageI, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongo-db:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database("Budgeting")

	accountS := NewAccount(db)
	budgetS := NewBudget(db)
	categoryS := NewCategory(db)
	goalS := NewGoal(db)
	transaction := NewTransaction(db)

	return &Storage{
		DB:        db,
		AccountS: accountS,
		BudgetS:  budgetS,
		CategoryS: categoryS,
		GoalS: goalS,
		TransactionS: transaction,
	}, nil
}

func (s *Storage) Account() storage.AccountI {
	if s.AccountS == nil {
		s.AccountS = NewAccount(s.DB)
	}
	return s.AccountS
}

func (s *Storage) Budget() storage.BudgetI {
	if s.BudgetS == nil {
		s.BudgetS = NewBudget(s.DB)
	}
	return s.BudgetS
}

func (s *Storage) Category() storage.CategoryI{
	if s.CategoryS == nil{
		s.CategoryS = NewCategory(s.DB)
	}
	return s.CategoryS
}

func (s *Storage) Goal() storage.GoalI{
	if s.GoalS == nil{
		s.GoalS = NewGoal(s.DB)
	}
	return s.GoalS
}

func (s *Storage) Transaction() storage.TransactionI{
	if s.TransactionS == nil{
		s.TransactionS = NewTransaction(s.DB)
	}
	return s.TransactionS
}