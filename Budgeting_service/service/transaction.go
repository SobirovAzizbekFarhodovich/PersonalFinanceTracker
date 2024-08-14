package service

import (
	"context"
	pb "budgeting/genprotos"
	"budgeting/storage"

	"github.com/google/uuid"
)

type TransactionService struct {
	stg storage.StorageI
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(stg *storage.StorageI) *TransactionService {
	return &TransactionService{stg: *stg}
}

func (s *TransactionService) CreateTransaction(c context.Context,req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	id := uuid.NewString()
	req.Transaction.Id = id
	_, err := s.stg.Transaction().CreateTransaction(req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTransactionResponse{}, nil
}

func (s *TransactionService) UpdateTransaction(c context.Context,req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	_, err := s.stg.Transaction().UpdateTransaction(req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTransactionResponse{}, nil
}

func (s *TransactionService) DeleteTransaction(c context.Context,req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	_, err := s.stg.Transaction().DeleteTransaction(req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTransactionResponse{}, nil
}

func (s *TransactionService) GetTransaction(c context.Context,req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	res, err := s.stg.Transaction().GetTransaction(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) ListTransactions(c context.Context,req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	res, err := s.stg.Transaction().ListTransactions(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) Spending(c context.Context,req *pb.SpendingRequest) (*pb.SpendingResponse, error){
	res, err := s.stg.Transaction().Spending(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) Income(c context.Context,req *pb.IncomeRequest) (*pb.IncomeResponse, error){
	res, err := s.stg.Transaction().Income(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}