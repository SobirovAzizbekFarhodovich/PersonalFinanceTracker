package service

import (
	"context"
	pb "budgeting/genprotos"
	"budgeting/storage"

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
