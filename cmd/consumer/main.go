package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(configMap)

	if err != nil {
		fmt.Println("Error on create consumer: ", err.Error())
	}

	topics := []string{"transactions"}

	consumer.SubscribeTopics(topics, nil)

	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(message.Value), message.TopicPartition)
		}
	}
}
