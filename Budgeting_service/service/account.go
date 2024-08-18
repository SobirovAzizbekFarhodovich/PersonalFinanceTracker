package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type AccountService struct {
	stg storage.StorageI
	pb.UnimplementedAccountServiceServer
}

func NewAccountService(stg *storage.StorageI) *AccountService {
	return &AccountService{stg: *stg}
}

func (s *AccountService) CreateAccount(c context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if _, err := uuid.Parse(req.Account.UserId); err != nil {
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}
	id := uuid.NewString()
	req.Account.Id = id

	_, err := s.stg.Account().CreateAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{}, nil
}

func (s *AccountService) UpdateAccount(c context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	if _, err := uuid.Parse(req.Account.UserId); err != nil {
		return nil, fmt.Errorf("invalid UserId: must be a valid UUID")
	}
	_, err := s.stg.Account().UpdateAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateAccountResponse{}, nil
}

func (s *AccountService) DeleteAccount(c context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	_, err := s.stg.Account().DeleteAccount(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteAccountResponse{}, nil
}

func (s *AccountService) GetAccount(c context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	res, err := s.stg.Account().GetAccount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountService) ListAccounts(c context.Context, req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error) {
	res, err := s.stg.Account().ListAccounts(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountService) GetAmount(c context.Context, req *pb.GetAmountRequest) (*pb.GetAmountResponse, error) {
	res, err := s.stg.Account().GetAmount(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *AccountService) UpdateAmount(c context.Context, req *pb.UpdateAmountRequest) (*pb.UpdateAmountResponse, error) {
	_, err := s.stg.Account().UpdateAmount(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
