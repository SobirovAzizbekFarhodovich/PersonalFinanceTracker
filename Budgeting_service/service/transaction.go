package service

import (
	pb "budgeting/genprotos"
	"budgeting/storage"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type TransactionService struct {
	stg storage.StorageI
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(stg *storage.StorageI) *TransactionService {
	return &TransactionService{stg: *stg}
}

func (s *TransactionService) CreateTransaction(c context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {
	id := uuid.NewString()
	req.Transaction.Id = id
	_, err := s.stg.Transaction().CreateTransaction(req)
	if err != nil {
		return nil, err
	}

	Amountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: req.Transaction.UserId})
	if err != nil {
		return nil, err
	}

	if Amountres.Balance < req.Transaction.Amount && req.Transaction.Type == "expense" {
		return nil, errors.New("your balance is not more than the requested amount")
	}
	amount := Amountres.Balance
	if req.Transaction.Type == "expense" {
		amount -= req.Transaction.Amount
	} else if req.Transaction.Type == "income" {
		amount += req.Transaction.Amount
	}

	_, err = s.stg.Account().UpdateAmount(&pb.UpdateAmountRequest{UserId: req.Transaction.UserId, Balance: amount})
	if err != nil {
		return nil, errors.New("failed to update amount")
	}
	message := fmt.Sprintf("%f sum has been %s your account", req.Transaction.Amount, req.Transaction.Type)

	notification := &pb.CreateNotificationRequest{}
	notification.Notification = &pb.Notification{
		Id:      uuid.NewString(),
		UserId:  req.Transaction.UserId,
		Message: message,
	}

	_, err = s.stg.Notification().CreateNotification(&pb.CreateNotificationRequest{Notification: notification.Notification})
	if err != nil {
		return nil, errors.New("notification not created")
	}

	return &pb.CreateTransactionResponse{}, nil
}

func (s *TransactionService) UpdateTransaction(c context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	_, err := s.stg.Transaction().UpdateTransaction(req)
	if err != nil {
		return nil, err
	}

	Amountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: req.Transaction.UserId})
	if err != nil {
		return nil, errors.New("user not found")
	}

	Transactionres, err := s.stg.Transaction().GetTransaction(&pb.GetTransactionRequest{Id: req.Transaction.Id})
	if err != nil {
		return nil, errors.New("transaction not found")
	}

	amount := Amountres.Balance
	if Transactionres.Transaction.Type == "expense" {
		amount += Transactionres.Transaction.Amount
	} else if Transactionres.Transaction.Type == "income" {
		amount -= Transactionres.Transaction.Amount
	}

	if req.Transaction.Type == "expense" {
		amount -= req.Transaction.Amount
	} else if req.Transaction.Type == "income" {
		amount += req.Transaction.Amount
	}

	_, err = s.stg.Account().UpdateAmount(&pb.UpdateAmountRequest{UserId: req.Transaction.UserId, Balance: amount})
	if err != nil {
		return nil, errors.New("failed to update amount")
	}

	_, err = s.stg.Transaction().UpdateTransaction(&pb.UpdateTransactionRequest{Transaction: req.Transaction})
	if err != nil {
		return nil, errors.New("failed to update transaction")
	}
	message := fmt.Sprintf("%f sum has been updated in your account", req.Transaction.Amount)
	notification := &pb.CreateNotificationRequest{}
	notification.Notification = &pb.Notification{
		Id:      uuid.NewString(),
		UserId:  req.Transaction.UserId,
		Message: message,
	}

	_, err = s.stg.Notification().CreateNotification(&pb.CreateNotificationRequest{Notification: notification.Notification})
	if err != nil {
		return nil, errors.New("notification not created")
	}

	return &pb.UpdateTransactionResponse{}, nil
}

func (s *TransactionService) DeleteTransaction(c context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	
	gettr, err := s.stg.Transaction().GetTransaction(&pb.GetTransactionRequest{Id: req.Id})
	if err != nil {
		return nil, errors.New("transaction not found")
	}
	Amountres, err := s.stg.Account().GetAmount(&pb.GetAmountRequest{UserId: gettr.Transaction.UserId})
	if err != nil {
		return nil, errors.New("user not found")
	}
	
	amount := Amountres.Balance
	if gettr.Transaction.Type == "expense" {
		amount += gettr.Transaction.Amount
	} else if gettr.Transaction.Type == "income" {
		amount -= gettr.Transaction.Amount
	}
	
	_, err = s.stg.Account().UpdateAmount(&pb.UpdateAmountRequest{UserId: gettr.Transaction.UserId, Balance: amount})
	if err != nil {
		return nil, errors.New("failed to update amount")
	}

	message := fmt.Sprintf("%f sum has been reverted from your account due to transaction deletion", gettr.Transaction.Amount)
	notification := &pb.CreateNotificationRequest{}
	notification.Notification = &pb.Notification{
		Id:      uuid.NewString(),
		UserId:  gettr.Transaction.UserId,
		Message: message,
	}

	_, err = s.stg.Notification().CreateNotification(&pb.CreateNotificationRequest{Notification: notification.Notification})
	if err != nil {
		return nil, errors.New("notification not created")
	}

	_, err = s.stg.Transaction().DeleteTransaction(req)
	if err != nil {
		return nil, err
	}
	
	return &pb.DeleteTransactionResponse{}, nil
}

func (s *TransactionService) GetTransaction(c context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	res, err := s.stg.Transaction().GetTransaction(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) ListTransactions(c context.Context, req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	res, err := s.stg.Transaction().ListTransactions(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) Spending(c context.Context, req *pb.SpendingRequest) (*pb.SpendingResponse, error) {
	res, err := s.stg.Transaction().Spending(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TransactionService) Income(c context.Context, req *pb.IncomeRequest) (*pb.IncomeResponse, error) {
	res, err := s.stg.Transaction().Income(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
