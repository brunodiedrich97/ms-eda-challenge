package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com.br/brunodiedrich97/ms-balance/internal/database"
	"github.com.br/brunodiedrich97/ms-balance/internal/usecase/get_balance"
	"github.com.br/brunodiedrich97/ms-balance/internal/usecase/update_balance"
	"github.com.br/brunodiedrich97/ms-balance/internal/web"
	"github.com.br/brunodiedrich97/ms-balance/internal/web/webserver"
	"github.com.br/brunodiedrich97/ms-balance/pkg/kafka"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "db-balance", "3306", "balances"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	balanceDB := database.NewBalanceDB(db)
	updateBalanceUseCase := update_balance.NewUpdateBalanceUseCase(balanceDB)
	getBalanceUseCase := get_balance.NewGetBalanceUseCase(balanceDB)

	// Configuração do Kafka Consumer
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "balances",
		"auto.offset.reset": "earliest",
	}
	topics := []string{"balances"}
	kafkaConsumer := kafka.NewKakfaConsumer(configMap, topics)

	// Canal para receber mensagens do Kafka
	msgChan := make(chan *ckafka.Message)

	// Roda o Consumidor em background (Goroutine)
	go kafkaConsumer.Consume(msgChan)

	go func() {
		for msg := range msgChan {
			var kafkaMessage struct {
				Name    string                               `json:"Name"`
				Payload update_balance.UpdateBalanceInputDTO `json:"Payload"`
			}

			err := json.Unmarshal(msg.Value, &kafkaMessage)
			if err != nil {
				fmt.Println("Error unmarshalling kafka message:", err)
				continue
			}

			input := kafkaMessage.Payload

			if input.AccountIDFrom != "" {
				updateBalanceUseCase.Execute(input)
				fmt.Printf("Balance updated for accounts: %s and %s\n", input.AccountIDFrom, input.AccountIDTo)
			} else {
				fmt.Println("Received message with empty payload:", string(msg.Value))
			}
		}
	}()

	webserver := webserver.NewWebServer(":3003")
	balanceHandler := web.NewWebBalanceHandler(*getBalanceUseCase)

	webserver.AddHandler("/balances/{account_id}", balanceHandler.GetBalance)

	fmt.Println("Server is running on port 3003")
	webserver.Start()
}
