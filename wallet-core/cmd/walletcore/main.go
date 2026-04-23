package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com.br/brunodiedrich97/ms-wallet/internal/database"
	"github.com.br/brunodiedrich97/ms-wallet/internal/event"
	"github.com.br/brunodiedrich97/ms-wallet/internal/event/handler"
	"github.com.br/brunodiedrich97/ms-wallet/internal/usecase/create_account"
	"github.com.br/brunodiedrich97/ms-wallet/internal/usecase/create_client"
	"github.com.br/brunodiedrich97/ms-wallet/internal/usecase/create_transaction"
	"github.com.br/brunodiedrich97/ms-wallet/internal/web"
	"github.com.br/brunodiedrich97/ms-wallet/internal/web/webserver"
	"github.com.br/brunodiedrich97/ms-wallet/pkg/events"
	"github.com.br/brunodiedrich97/ms-wallet/pkg/kafka"
	"github.com.br/brunodiedrich97/ms-wallet/pkg/uow"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "db-wallet", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	configMag := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMag)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewUpdateBalanceKafkaHandler(kafkaProducer))
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)

	// Uow - Unit of Work
	ctx := context.Background()
	uow := uow.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	// Instanciando na mão, sem container de DI
	createClientUseCase := create_client.NewCreateClientUseCase(clientDB)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDB)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webserver := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webserver.AddHandler("/clients", clientHandler.CreateClient)
	webserver.AddHandler("/accounts", accountHandler.CreateAccount)
	webserver.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running on port 8080")
	webserver.Start()
}
