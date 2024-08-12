package extra

import (
	pb "budgeting/genprotos"
	m "budgeting/models"
)

func AccountToBSON(p *pb.Account) *m.Account {
	return &m.Account{
		ID:       p.Id,
		UserID:   p.UserId,
		Name:     p.Name,
		Type:     p.Type,
		Balance:  p.Balance,
		Currency: p.Currency,
	}
}

func BsonToAccount(b *m.Account) *pb.Account {
	return &pb.Account{
		Id:       b.ID,
		UserId:   b.UserID,
		Name:     b.Name,
		Type:     b.Type,
		Balance:  b.Balance,
		Currency: b.Currency,
	}
}

func CategoryToBSON(p *pb.Category) *m.Category {
	return &m.Category{
		ID:     p.Id,
		UserID: p.UserId,
		Name:   p.Name,
		Type:   p.Type,
	}
}

func BsonToCategory(b *m.Category) *pb.Category {
	return &pb.Category{
		Id:     b.ID,
		UserId: b.UserID,
		Name:   b.Name,
		Type:   b.Type,
	}
}

func BudgetToBSON(p *pb.Budget) *m.Budget {
	return &m.Budget{
		ID:         p.Id,
		UserID:     p.UserId,
		CategoryID: p.CategoryId,
		Period:     p.Period,
		Amount:     p.Amount,
		StartDate:  p.StartDate,
		EndDate:    p.EndDate,
	}
}

func BsonToBudget(b *m.Budget) *pb.Budget {
	return &pb.Budget{
		Id:         b.ID,
		UserId:     b.UserID,
		CategoryId: b.CategoryID,
		Period:     b.Period,
		Amount:     b.Amount,
		StartDate:  b.StartDate,
		EndDate:    b.EndDate,
	}
}

func GoalToBSON(p *pb.Goal) *m.Goal {
	return &m.Goal{
		ID:            p.Id,
		UserID:        p.UserId,
		Name:          p.Name,
		TargetAmount:  p.TargetAmount,
		CurrentAmount: p.CurrentAmount,
		Deadline:      p.Deadline,
		Status:        p.Status,
	}
}

func BsonToGoal(b *m.Goal) *pb.Goal {
	return &pb.Goal{
		Id:            b.ID,
		UserId:        b.UserID,
		Name:          b.Name,
		TargetAmount:  b.TargetAmount,
		CurrentAmount: b.CurrentAmount,
		Deadline:      b.Deadline,
		Status:        b.Status,
	}
}

func TransactionToBSON(p *pb.Transaction) *m.Transaction {
	return &m.Transaction{
		ID:          p.Id,
		UserID:      p.UserId,
		CategoryID:  p.CategoryId,
		AccountID:   p.AccountId,
		Amount:      p.Amount,
		Type:        p.Type,
		Description: p.Description,
		Date:        p.Date,
	}
}

func BsonToTransaction(b *m.Transaction) *pb.Transaction {
	return &pb.Transaction{
		Id:          b.ID,
		UserId:      b.UserID,
		CategoryId:  b.CategoryID,
		AccountId:   b.AccountID,
		Amount:      b.Amount,
		Type:        b.Type,
		Description: b.Description,
		Date:        b.Date,
	}
}
