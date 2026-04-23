package database

import (
	"database/sql"

	"github.com.br/brunodiedrich97/ms-balance/internal/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{DB: db}
}

func (b *BalanceDB) GetByAccountID(accountID string) (*entity.Balance, error) {
	var balance entity.Balance
	stmt, err := b.DB.Prepare("SELECT id, account_id, amount, created_at FROM balances WHERE account_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(accountID)
	err = row.Scan(&balance.ID, &balance.AccountID, &balance.Amount)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

// Save executa a lógica de Upsert.
// No contexto de EDA, se o microsserviço de Balances ainda não conhece a conta,
// ele a cria com o valor da primeira transação que capturar.
func (b *BalanceDB) Save(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare(`
		INSERT INTO balances (id, account_id, amount) 
		VALUES (?, ?, ?) 
		ON DUPLICATE KEY UPDATE amount = values(amount)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(balance.ID, balance.AccountID, balance.Amount)
	if err != nil {
		return err
	}
	return nil
}
