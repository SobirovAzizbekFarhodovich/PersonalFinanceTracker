package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"

	"github.com/google/uuid"
)

type NotificationService struct {
	stg storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(stg *storage.StorageI) *NotificationService {
	return &NotificationService{stg: *stg}
}

func (s *NotificationService) CreateNotification(c context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	id := uuid.NewString()
	req.Notification.Id = id
	_, err := s.stg.Notification().CreateNotification(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateNotificationResponse{}, nil
}

func (s *NotificationService) GetNotification(c context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	res, err := s.stg.Notification().GetNotification(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
