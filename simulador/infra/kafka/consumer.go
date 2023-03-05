package kafka

import (
	"fmt"
	"log"
	"os"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConsumer struct {
	MsgChan chan *ckafka.Message // canal do GO que recebe todas as mensagens do kafka
}

func (k *KafkaConsumer) Consume() {

	//definindo configmap do kafka
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Erro ao consumir menssagem do Kafka" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopics")}
	c.SubscribeTopics(topics, nil) // ao se inscrever o kafka estará observando e consumindo as msg do topico informado

	fmt.Println("Kafka Consumer foi iniciado!")

	for {
		msg, err := c.ReadMessage(-1) // -1 é o timeout para não ficar esperando pelas menssagens
		if err == nil {               // se não receber nunhum erro, então pega essa menssagem e joga no canal do GO
			k.MsgChan <- msg
		}
	}

}
