package kafka


import (
	"log/slog"
	"os"
)

func InitKafkaConsumer(){
	if(consumer == nil){
		err := createConsumer()
		if err != nil {
			slog.Error("Failed to create consumer")
			os.Exit(1)
		}
	}
}

func CloseKafkaConsumer(){
	consumer.Close()
}