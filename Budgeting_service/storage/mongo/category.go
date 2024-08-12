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

type Category struct {
	mongo *mongo.Collection
}

func NewCategory(db *mongo.Database) *Category {
	return &Category{mongo: db.Collection("category")}
}

func (s *Category) CreateCategory(req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	bsonAccount := e.CategoryToBSON(req.Category)
	_, err := s.mongo.InsertOne(context.TODO(), bsonAccount)
	if err != nil {
		return nil, err

	}
	return &pb.CreateCategoryResponse{}, nil
}

func (s *Category) UpdateCategory(req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	bsonAccount := e.CategoryToBSON(req.Category)
	filter := bson.M{"_id": bsonAccount.ID}
	update := bson.M{"$set": bsonAccount}

	_, err := s.mongo.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCategoryResponse{}, nil
}

func (s *Category) DeleteCategory(req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	_, err := s.mongo.DeleteOne(context.TODO(), bson.M{"_id": req.Id})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryResponse{}, nil
}

func (s *Category) GetCategory(req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	var bsonAccount m.Category
	err := s.mongo.FindOne(context.TODO(), bson.M{"_id": req.Id}).Decode(&bsonAccount)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("progress not found")
	} else if err != nil {
		return nil, err
	}
	return &pb.GetCategoryResponse{Category: e.BsonToCategory(&bsonAccount)}, nil
}

func (s *Category) ListCategories(req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
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

		}},
	}

	cursor, err := s.mongo.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Category
	for cursor.Next(context.TODO()) {
		var account pb.Category
		if err := cursor.Decode(&account); err != nil {
			return nil, err
		}
		accounts = append(accounts, &account)
	}

	return &pb.ListCategoriesResponse{
		Categories: accounts,
	}, nil
}
