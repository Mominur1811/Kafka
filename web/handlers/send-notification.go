package handlers

import (
	"kafka-by-momin/kafka"
	"net/http"
)

func CreateNotification(w http.ResponseWriter, r *http.Request) {

	kafka.PublishMessage("my-topic", "Hi this is momin")
}
