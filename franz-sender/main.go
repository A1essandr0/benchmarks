package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/twmb/franz-go/pkg/kgo"
)


func main() {
	run(config)
}


func run(config Config) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(config.brokers...),
		kgo.ConsumerGroup(config.consumerGroupName),
		kgo.ConsumeTopics(config.sourceTopic),
	)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan bool)
	log.Println("Application started")

	go func() {
		defer close(done)
		for {
			fetches := client.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				fmt.Println(errs)
				break
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				incomingRecord := iter.Next()

				msg := RawEvent{}
				rawmsg := incomingRecord.Value
				err := json.Unmarshal(rawmsg, &msg)
				log.Printf("Received raw: %+v", rawmsg)
				if err != nil {
					log.Printf("Error deserializing: %+v", rawmsg)
					continue
				}
				log.Printf("Deserialized: %+v", msg)
				
				// producing result
				marshalled, err := json.Marshal(msg)
				if err != nil {
					log.Printf("error marshalling data %v; %v", msg, err)
					continue
				}
				log.Println("Sending serialized further: ", string(marshalled))
				record := &kgo.Record{Topic: config.destinationTopic, Value: marshalled}
				if err := client.ProduceSync(ctx, record).FirstErr(); err != nil {
					fmt.Printf("record had a produce error while synchronously producing: %v\n", err)
				}

			}
		}
	}()

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	<- wait
	log.Printf("Gracefully stoping...")
	cancel()
	<- done
	log.Printf("...gracefully stopped")
}