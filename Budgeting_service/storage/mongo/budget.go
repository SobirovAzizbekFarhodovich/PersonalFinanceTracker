package mongo

import (
	"context"

	e "budgeting/extra"
	pb "budgeting/genprotos"

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
	var budget pb.Budget
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
	return &pb.GetBudgetResponse{
		Budget:     &budget,
		Categories: &category,
	}, nil
}

func (s *Budget) ListBudgets(req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
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

	var budgets []*pb.GetBudgetResponse
	for cursor.Next(context.TODO()) {
		var budget pb.Budget
		if err := cursor.Decode(&budget); err != nil {
			return nil, err
		}
		var category pb.Category
		categoryFilter := bson.M{"category_id": budget.CategoryId}
		err := s.mongo.FindOne(context.TODO(), categoryFilter).Decode(&category)
		if err != nil {
			return nil, err
		}

		budgets = append(budgets, &pb.GetBudgetResponse{
			Budget:    &budget,
			Categories: &category,
		})
	}
	return &pb.ListBudgetsResponse{
		Budgets: budgets,
	}, nil
}
