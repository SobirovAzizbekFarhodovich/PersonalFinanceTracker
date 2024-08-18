package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type CategoryService struct {
	stg storage.StorageI
	pb.UnimplementedCategoryServiceServer
}

func NewCategoryService(stg *storage.StorageI) *CategoryService {
	return &CategoryService{stg: *stg}
}

func (s *CategoryService) CreateCategory(c context.Context,req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	id := uuid.NewString()
	req.Category.Id = id

	if _, err := uuid.Parse(req.Category.UserId);err != nil{
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}

	if req.Category.Type != "expense" && req.Category.Type != "income" {
		return nil, fmt.Errorf("invalid Category Type: must be either 'expense' or 'income'")
	}

	_, err := s.stg.Category().CreateCategory(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCategoryResponse{}, nil
}

func (s *CategoryService) UpdateCategory(c context.Context,req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error) {
	_, err := s.stg.Category().UpdateCategory(req)
	if err != nil {
		return nil, err
	}
	if _, err := uuid.Parse(req.Category.UserId);err != nil{
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}

	if req.Category.Type != "expense" && req.Category.Type != "income" {
		return nil, fmt.Errorf("invalid Category Type: must be either 'expense' or 'income'")
	}
	return &pb.UpdateCategoryResponse{}, nil
}

func (s *CategoryService) DeleteCategory(c context.Context,req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error) {
	_, err := s.stg.Category().DeleteCategory(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCategoryResponse{}, nil
}

func (s *CategoryService) GetCategory(c context.Context,req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error) {
	res, err := s.stg.Category().GetCategory(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CategoryService) ListCategories(c context.Context,req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	res, err := s.stg.Category().ListCategories(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
