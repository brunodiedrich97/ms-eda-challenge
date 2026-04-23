package create_transaction

import (
	"context"
	"testing"

	"github.com.br/brunodiedrich97/ms-wallet/internal/entity"
	"github.com.br/brunodiedrich97/ms-wallet/internal/event"
	"github.com.br/brunodiedrich97/ms-wallet/internal/usecase/mocks"
	"github.com.br/brunodiedrich97/ms-wallet/pkg/events"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe I", "j1@j.com")
	accountFrom, _ := entity.NewAccount(client1)
	accountFrom.Credit(1000)

	client2, _ := entity.NewClient("Jane Doe II", "j2@j.com")
	accountTo, _ := entity.NewAccount(client2)
	accountTo.Credit(1000)

	mockUow := &mocks.UowMock{}
	mockUow.On("Do", mock.Anything, mock.Anything).Return(nil)

	dispatcher := events.NewEventDispatcher()
	eventTransaction := event.NewTransactionCreated()
	eventBalance := event.NewBalanceUpdated()
	ctx := context.Background()

	uc := NewCreateTransactionUseCase(mockUow, dispatcher, eventTransaction, eventBalance)

	input := &CreateTransactionInputDTO{
		AccountIDFrom: accountFrom.ID,
		AccountIDTo:   accountTo.ID,
		Amount:        100,
	}

	output, err := uc.Execute(ctx, input)

	if err != nil {
		t.Error(err)
	}

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockUow.AssertExpectations(t)
	mockUow.AssertNumberOfCalls(t, "Do", 1)
}
