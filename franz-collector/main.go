package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/twmb/franz-go/pkg/kgo"
)

var (
	client *kgo.Client
)

func main() {
	var err error
	client, err = kgo.NewClient(
		kgo.SeedBrokers(config.brokers...),
		kgo.ConsumerGroup(config.consumerGroupName),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

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

		record := &kgo.Record{Topic: config.sourceTopic, Value: marshalled}
		if err := client.ProduceSync(nil, record).FirstErr(); err != nil {
			log.Printf("record had a produce error while synchronously producing: %v\n", err)
		} else {
			log.Println("sent OK")
		}
	}
}
