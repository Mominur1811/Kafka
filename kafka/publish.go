package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Publisher struct {
	Message interface{}
	Topic   string
}

func PublishMessage(topic string, message string) {
	msg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}

	// Produce the message
	err := producer.Produce(msg, nil)
	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	} else {
		log.Printf("Produced message: %s", message)
	}
}
