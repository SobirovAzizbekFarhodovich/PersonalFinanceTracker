package models

type Account struct {
	ID       string  `bson:"_id,omitempty"`
	UserID   string  `bson:"user_id"`
	Name     string  `bson:"name"`
	Type     string  `bson:"type"`
	Balance  float32 `bson:"balance"`
	Currency string  `bson:"currency"`
}

type Category struct {
	ID     string `bson:"_id,omitempty"`
	UserID string `bson:"user_id"`
	Name   string `bson:"name"`
	Type   string `bson:"type"`
}

type Budget struct {
	ID         string  `bson:"_id,omitempty"`
	UserID     string  `bson:"user_id"`
	CategoryID string  `bson:"category_id"`
	Period     string  `bson:"period"`
	Amount     float32 `bson:"amount"`
	StartDate  string  `bson:"start_date"`
	EndDate    string  `bson:"end_date"`
}

type Goal struct {
	ID            string  `bson:"_id,omitempty"`
	UserID        string  `bson:"user_id"`
	Name          string  `bson:"name"`
	TargetAmount  float32 `bson:"target_amount"`
	CurrentAmount float32 `bson:"current_amount"`
	Deadline      string  `bson:"deadline"`
	Status        string  `bson:"status"`
}

type Transaction struct {
	ID          string  `bson:"_id,omitempty"`
	UserID      string  `bson:"user_id"`
	CategoryID  string  `bson:"category_id"`
	AccountID   string  `bson:"account_id"`
	Amount      float32 `bson:"amount"`
	Type        string  `bson:"type"`
	Description string  `bson:"description"`
	Date        string  `bson:"date"`
}

type Notification struct{
	ID string `bson:"_id,omitempty"`
	UserID string `bson:"user_id"`
	Message string `bson:"message"`
}