package storage

import (
	pb "budgeting/genprotos"
)

type StorageI interface {
	Account() AccountI
	Budget() BudgetI
	Category() CategoryI
	Goal() GoalI
	Transaction() TransactionI
	Notification() NotificationI
}

type AccountI interface{
	CreateAccount(req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error)
	UpdateAccount(req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error)
	DeleteAccount(req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error)
	GetAccount(req *pb.GetAccountRequest) (*pb.GetAccountResponse, error)
	ListAccounts(req *pb.ListAccountsRequest) (*pb.ListAccountsResponse, error)
	GetAmount(req *pb.GetAmountRequest) (*pb.GetAmountResponse, error)
	UpdateAmount(req *pb.UpdateAmountRequest) (*pb.UpdateAmountResponse, error)
}

type BudgetI interface{
	CreateBudget(req *pb.CreateBudgetRequest) (*pb.CreateBudgetResponse, error)
	UpdateBudget(req *pb.UpdateBudgetRequest) (*pb.UpdateBudgetResponse, error)
	DeleteBudget(req *pb.DeleteBudgetRequest) (*pb.DeleteBudgetResponse, error)
	GetBudget(req *pb.GetBudgetRequest) (*pb.GetBudgetResponse, error)
	ListBudgets(req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error)
	GenerateBudgetPerformanceReport(req *pb.GenerateBudgetPerformanceReportRequest) (*pb.GenerateBudgetPerformanceReportResponse, error)
}

type CategoryI interface{
	CreateCategory(req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error)
	UpdateCategory(req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryResponse, error)
	DeleteCategory(req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryResponse, error)
	GetCategory(req *pb.GetCategoryRequest) (*pb.GetCategoryResponse, error)
	ListCategories(req *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error)
}

type GoalI interface {
	CreateGoal(req *pb.CreateGoalRequest) (*pb.CreateGoalResponse, error)
	UpdateGoal(req *pb.UpdateGoalRequest) (*pb.UpdateGoalResponse, error)
	DeleteGoal(req *pb.DeleteGoalRequest) (*pb.DeleteGoalResponse, error)
	GetGoal(req *pb.GetGoalRequest) (*pb.GetGoalResponse, error)
	ListGoals(req *pb.ListGoalsRequest) (*pb.ListGoalsResponse, error)
	GenerateGoalProgressReport(req *pb.GenerateGoalProgressReportRequest) (*pb.GenerateGoalProgressReportResponse, error)

}

type TransactionI interface {
	CreateTransaction(req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error)
	UpdateTransaction(req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error)
	DeleteTransaction(req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error)
	GetTransaction(req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error)
	ListTransactions(req *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error)
	Spending(req *pb.SpendingRequest) (*pb.SpendingResponse, error)
	Income(req *pb.IncomeRequest) (*pb.IncomeResponse, error)
}

type NotificationI interface{
	CreateNotification(req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error)
	GetNotification(req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error)
}