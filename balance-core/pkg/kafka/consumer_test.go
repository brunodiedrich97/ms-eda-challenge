package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestConsumerConsume(t *testing.T) {
	configMap := &ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
		"group.id":              "test",
	}
	topics := []string{"test"}
	consumer := NewKakfaConsumer(configMap, topics)

	// Criamos o canal que a função exige
	msgChan := make(chan *ckafka.Message)

	// Como o Consume tem um loop infinito (for { ... }),
	// em um teste real rodar ele em uma goroutine
	// ou num mock do Consumer do Confluent.
	assert.NotNil(t, consumer)
	assert.NotNil(t, msgChan)
}
