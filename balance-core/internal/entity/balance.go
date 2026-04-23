package entity

import (
	"github.com/google/uuid"
)

type Balance struct {
	ID        string
	AccountID string
	Amount    float64
}

func NewBalance(accountID string, amount float64) *Balance {
	return &Balance{
		ID:        uuid.New().String(),
		AccountID: accountID,
		Amount:    amount,
	}
}
