package kafka

import (
	"log/slog"
	"os"
)

func InitKafkaProducer(){
	if(producer == nil){
		err := createProducer()
		if err != nil {
			slog.Error("Failed to create new producer")
			os.Exit(1)
		}
	}
}

func CloseKafkaProducer(){
	producer.Close()
}