package gateway

import "github.com.br/brunodiedrich97/ms-balance/internal/entity"

type BalanceGateway interface {
	Save(balance *entity.Balance) error
	GetByAccountID(accountID string) (*entity.Balance, error)
}
