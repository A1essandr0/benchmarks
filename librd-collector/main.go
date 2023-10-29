package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var (
	producer *kafka.Producer
)

func main() {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.broker,
	})
	if err != nil {
		log.Fatalf("Couldn't initialize producer: %+v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	server := &http.Server{
		Addr: config.port,
		Handler: mux,
	}

	log.Printf("Starting server on %s...", config.port)
	err = server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
			log.Println("...server closed")
	} else if err != nil {
			log.Printf("Error starting server: %s", err)
	}
}


func getRoot(w http.ResponseWriter, r *http.Request) {
	event := RawEvent{
			Source: r.URL.Query().Get("source"),
			EventName: r.URL.Query().Get("event_name"),
			EventStatus: r.URL.Query().Get("event_status"),
			Created: r.URL.Query().Get("created"),
			Payout: r.URL.Query().Get("payout"),
	}
	marshalled, err := json.Marshal(event)
	serialized := string(marshalled)

	if err == nil {
		io.WriteString(w, "OK")
		log.Printf("received OK: %s", serialized)

		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &config.sourceTopic, Partition: kafka.PartitionAny},
			Value: marshalled,
			}, nil,
		)
		if err != nil {
			log.Printf("Could not send message to kafka. Error: %s", err.Error())
		}
		log.Println("sent OK")
	}       
}