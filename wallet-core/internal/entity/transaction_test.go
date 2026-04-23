package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1.ID, transaction.AccountFrom.ID)
	assert.Equal(t, account2.ID, transaction.AccountTo.ID)
	assert.Equal(t, 100.0, transaction.Amount)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account1.Balance)
	assert.NotEmpty(t, transaction.ID)
	assert.NotEmpty(t, transaction.CreatedAt)
}

func TestNewTransactionWithoutBalance(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 2000.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient funds")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}

func TestNewTransactionNegativeAmount(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, -100.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}

func TestNewTransactionZeroAmount(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 0.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}

func TestNewTransactionAccountFromNil(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(nil, account2, 100.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "accountFrom and accountTo are required")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}

func TestNewTransactionAccountToNil(t *testing.T) {
	client1, _ := NewClient("John Doe I", "j1@j.com")
	account1, _ := NewAccount(client1)
	client2, _ := NewClient("Jane Doe II", "j2@j.com")
	account2, _ := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, nil, 100.0)
	assert.NotNil(t, err)
	assert.Error(t, err, "accountFrom and accountTo are required")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account1.Balance)
}
