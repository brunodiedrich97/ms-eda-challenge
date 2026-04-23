package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 0.0, account.Balance)
	assert.NotEmpty(t, account.ID)
	assert.NotEmpty(t, account.CreatedAt)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestCreateAccountWithInvalidClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.Error(t, err)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Credit(100.0)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, account.Balance)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestCreditAccountWithInvalidAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Credit(-100.0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	account.Credit(100.0)
	err := account.Debit(50.0)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, account.Balance)
	assert.NotEmpty(t, account.UpdatedAt)
}

func TestDebitAccountWithInvalidAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Debit(-100.0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, account.Balance)
}

func TestDebitAccountWithInsufficientBalance(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	err := account.Debit(100.0)
	assert.Error(t, err)
	assert.Equal(t, 0.0, account.Balance)
}
