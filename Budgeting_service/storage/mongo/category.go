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

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Category
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Category
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToCategory(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.ListCategoriesResponse{
		Categories: accounts,
	}, nil
}
