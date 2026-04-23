package update_balance

import (
	"github.com.br/brunodiedrich97/ms-balance/internal/entity"
	"github.com.br/brunodiedrich97/ms-balance/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(g gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: g,
	}
}

func (u *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) error {
	// Atualiza conta de origem
	accountFrom := entity.NewBalance(input.AccountIDFrom, input.BalanceAccountIDFrom)
	err := u.BalanceGateway.Save(accountFrom)
	if err != nil {
		return err
	}

	// Atualiza conta de destino
	accountTo := entity.NewBalance(input.AccountIDTo, input.BalanceAccountIDTo)
	err = u.BalanceGateway.Save(accountTo)
	if err != nil {
		return err
	}

	return nil
}
