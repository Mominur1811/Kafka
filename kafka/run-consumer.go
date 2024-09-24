package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Consume starts a goroutine that listens for messages
func Consume() {
	// Subscribe to the specified topic
	if err := consumer.SubscribeTopics([]string{"my-topic"}, nil); err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	// Start the message consuming loop
	go tryConsume(GetConsumer())
}

// tryConsume continuously reads messages from the Kafka topic
func tryConsume(consumer *kafka.Consumer) {
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			// Handle error reading message
			if kafkaErr, ok := err.(kafka.Error); ok {
				if kafkaErr.IsFatal() {
					log.Fatalf("Fatal error: %s", kafkaErr)
				}
				log.Printf("Error reading message: %s\n", err)
			}
			continue
		}

		// Successfully read a message
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	}
}
