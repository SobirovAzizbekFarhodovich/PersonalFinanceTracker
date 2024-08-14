package mongo

import (
	"context"

	e "budgeting/extra"
	pb "budgeting/genprotos"
	m "budgeting/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Notification struct {
	mongo *mongo.Collection
}

func NewNotification(db *mongo.Database) *Notification {
	return &Notification{mongo: db.Collection("Notification")}
}

func (s *Notification) CreateNotification(req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	bsonNotification := e.NotificationToBson(req.Notification)
	_, err := s.mongo.InsertOne(context.TODO(), bsonNotification)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotificationResponse{}, nil
}

func (s *Notification) GetNotification(req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	limit := req.Limit
	page := req.Page
	skip := (page - 1) * limit

	cursor, err := s.mongo.Find(context.TODO(), bson.M{}, options.Find().SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var accounts []*pb.Notification
	for cursor.Next(context.TODO()) {
		var bsonAccount m.Notification
		if err := cursor.Decode(&bsonAccount); err != nil {
			return nil, err
		}
		account := e.BsonToNotification(&bsonAccount)
		accounts = append(accounts, account)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.GetNotificationResponse{
		Notification: accounts,
	}, nil
}
