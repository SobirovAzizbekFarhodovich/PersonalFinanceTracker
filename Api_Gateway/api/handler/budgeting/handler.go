package handler

import (
	budgeting "api/genprotos/budgeting"

	"google.golang.org/grpc"
)

type BudgetingHandler struct {
	Account      budgeting.AccountServiceClient
	Budget       budgeting.BudgetServiceClient
	Category     budgeting.CategoryServiceClient
	Goal         budgeting.GoalServiceClient
	Transaction  budgeting.TransactionServiceClient
	Notification budgeting.NotificationServiceClient
	Redis        InMemoryStorageI
}

func NewBudgetingHandler(budgetingConn *grpc.ClientConn, storageRedis InMemoryStorageI) *BudgetingHandler {
	return &BudgetingHandler{
		Account:      budgeting.NewAccountServiceClient(budgetingConn),
		Budget:       budgeting.NewBudgetServiceClient(budgetingConn),
		Category:     budgeting.NewCategoryServiceClient(budgetingConn),
		Goal:         budgeting.NewGoalServiceClient(budgetingConn),
		Transaction:  budgeting.NewTransactionServiceClient(budgetingConn),
		Notification: budgeting.NewNotificationServiceClient(budgetingConn),
		Redis:        storageRedis,
	}
}
