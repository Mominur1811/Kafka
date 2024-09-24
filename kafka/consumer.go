package kafka

import (
	"kafka-by-momin/config"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var consumer *kafka.Consumer

func createConsumer() error{
	config := &kafka.ConfigMap{
		"bootstrap.servers": config.GetConfig().Kafka.Consumer.Broker,
		"group.id":          config.GetConfig().Kafka.Consumer.GroupId,
		"auto.offset.reset": config.GetConfig().Kafka.Consumer.AutoOffsetReset, // Start reading from the earliest message
	}

	var err error
	consumer, err = kafka.NewConsumer(config)
	return err
}

func GetConsumer()*kafka.Consumer {
	if(consumer == nil){
		InitKafkaConsumer()
	}

	return consumer
}