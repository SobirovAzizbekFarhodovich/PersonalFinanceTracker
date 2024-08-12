package service

import (
	"context"
	pb "budgeting/genprotos"
	"budgeting/storage"

	"github.com/google/uuid"
)

type AccountService struct {
	stg storage.StorageI
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(stg *storage.StorageI) *AccountService {
	return &AccountService{stg: *stg}
}

func (s *AccountService) CreateAccount(c context.Context,req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	id := uuid.NewString()
	req.Account.Id = id
	_, err := s.stg.Account().CreateAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{}, nil
}

func (s *AccountService) UpdateAccount(c context.Context,req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	_, err := s.stg.Account().UpdateAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountResponse{}, nil
}

func (s *AccountService) DeleteAccount(c context.Context,req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	_, err := s.stg.Account().DeleteAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountResponse{}, nil
}

func (s *AccountService) GetAccount(c context.Context,req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	res, err := s.stg.Account().GetAccount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountService) ListAccounts(c context.Context,req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	res, err := s.stg.Account().ListAccounts(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
