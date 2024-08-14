package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Budget struct {
	mongo *mongo.Collection
}

func NewBudget(db *mongo.Database) *Budget {
	return &Budget{mongo: db.Collection("budget")}
}

func (s *Budget) CreateBudget(req *pb.CreateBudgetRequest) (*pb.CreateBudgetResponse, error) {
	bsonBudget := e.BudgetToBSON(req.Budget)
	_, err := s.mongo.InsertOne(context.TODO(), bsonBudget)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBudgetResponse{}, nil
}

func (s *Budget) UpdateBudget(req *pb.UpdateBudgetRequest) (*pb.UpdateBudgetResponse, error) {
	bsonBudget := e.BudgetToBSON(req.Budget)
	filter := bson.M{"_id": bsonBudget.ID}
	update := bson.M{"$set": bsonBudget}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateBudgetResponse{}, nil
}

func (s *Budget) DeleteBudget(req *pb.DeleteBudgetRequest) (*pb.DeleteBudgetResponse, error) {
	_, err := s.mongo.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBudgetResponse{}, nil

}

func (s *Budget) GetBudget(req *pb.GetBudgetRequest) (*pb.GetBudgetResponse, error) {
	var bsonGoal m.Budget
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonGoal)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("budget not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetBudgetResponse{Budget: e.BsonToBudget(&bsonGoal)}, nil
}

func (s *Budget) ListBudgets(req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Budget
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Budget
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToBudget(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListBudgetsResponse{
		Budgets: accounts,
	}, nil
}

func (s *Budget) GenerateBudgetPerformanceReport(req *pb.GenerateBudgetPerformanceReportRequest) (*pb.GenerateBudgetPerformanceReportResponse, error) {
	var bsonBudget m.Budget
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonBudget)
	if err == mongo.ErrNoDocuments {
		return nil, fmt.Errorf("budget not found")
	} else if err != nil {
		return nil, err
	}
	budget := e.BsonToBudget(&bsonBudget)

	startTime, err := time.Parse(time.RFC3339, budget.StartDate+"T00:00:00Z")
	if err != nil {
		return nil, fmt.Errorf("invalid start date format")
	}

	endTime, err := time.Parse(time.RFC3339, budget.EndDate+"T23:59:59Z")
	if err != nil {
		return nil, fmt.Errorf("invalid end date format")
	}

	valid := false
	switch budget.Period {
	case "daily":
		valid = endTime.Sub(startTime).Hours() <= 24
	case "weekly":
		valid = endTime.Sub(startTime).Hours() <= 7*24
	case "monthly":
		valid = endTime.Sub(startTime).Hours() <= 31*24
	case "yearly":
		valid = endTime.Sub(startTime).Hours() <= 365*24
	default:
		return nil, fmt.Errorf("invalid period type")
	}

	if !valid {
		return nil, fmt.Errorf("start_date and end_date do not match the selected period")
	}

	spentamount := float32(0)

	resp := &pb.GenerateBudgetPerformanceReportResponse{
		Id:         budget.Id,
		UserId:     budget.UserId,
		CategoryId: budget.CategoryId,
		Amount:     budget.Amount,
		Period:     budget.Period,
		StartDate:  budget.StartDate,
		EndDate:    budget.EndDate,
		SpentAmount: spentamount,
	}

	return resp, nil
}
