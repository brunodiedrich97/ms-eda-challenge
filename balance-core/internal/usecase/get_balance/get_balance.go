package get_balance

import "github.com.br/brunodiedrich97/ms-balance/internal/gateway"

type GetBalanceOutputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(g gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: g,
	}
}

func (u *GetBalanceUseCase) Execute(accountID string) (*GetBalanceOutputDTO, error) {
	// Busca o saldo consolidado no banco de dados do microsserviço Balances
	balance, err := u.BalanceGateway.GetByAccountID(accountID)
	if err != nil {
		return nil, err
	}

	return &GetBalanceOutputDTO{
		AccountID: balance.AccountID,
		Balance:   balance.Amount,
	}, nil
}
