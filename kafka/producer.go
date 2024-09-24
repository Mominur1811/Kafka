package kafka

import (
	"kafka-by-momin/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func createProducer()error{
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().Kafka.Producer.Broker,
	})
	
	return err
}

func GetProducer()*kafka.Producer{
	InitKafkaProducer()
	return producer
}