package gateway

import "github.com.br/brunodiedrich97/ms-wallet/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
