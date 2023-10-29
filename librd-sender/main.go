package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	kafkaConsumerConfig := &kafka.ConfigMap{
		"bootstrap.servers": config.broker,
		"group.id": config.consumerGroupName,
		// "security.protocol": conf.KafkaSecurityProtocol,
		// "sasl.mechanism": conf.KafkaSaslMechanism,
		// "sasl.username": conf.KafkaUsername,
		// "sasl.password": conf.KafkaPassword,
		"auto.offset.reset": "earliest",
		"debug": "consumer",
	}

	consumer, err := kafka.NewConsumer(kafkaConsumerConfig)
	if err != nil {
		log.Fatalf("Couldn't initialize consumer: %+v", err)
	}
	err = consumer.SubscribeTopics([]string{config.sourceTopic}, nil)
	if err != nil {
		log.Fatalf("Couldn't subscribe to topic %s, consumer: %+v", config.sourceTopic, err)
	}

	kafkaProducerConfig := &kafka.ConfigMap{"bootstrap.servers": config.broker}
	producer, err := kafka.NewProducer(kafkaProducerConfig)
	if err != nil {
		log.Fatalf("Couldn't initialize producer: %+v", err)
	}

	exitSignalChannel := make(chan os.Signal, 1)
	signal.Notify(exitSignalChannel, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			msg, err := consumer.ReadMessage(time.Second)
			if err == nil {
				var event RawEvent
				err = json.Unmarshal(msg.Value, &event)
				if err != nil {
					log.Printf("Could not unmarshal event. Error: %s", err)
					continue
				}

				log.Println("received: ", event)
				marshalled, err := json.Marshal(event)
				if err != nil {
					log.Printf("Could not marshal event. Error: %s", err)
					continue

				}

				err = producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &config.destinationTopic, Partition: kafka.PartitionAny},
					Value: marshalled,
					},
					nil,
				)
				if err != nil {
					log.Printf("Could not send message to kafka. Error: %s", err.Error())
					continue
				}
				log.Println("sent further OK")

			} else if !err.(kafka.Error).IsTimeout() {
				if msg != nil {
					log.Printf("Partition-specific error: %+v (%+v)\n", err, msg.TopicPartition)
				}
				log.Printf("Consumer error: %+v (%+v)\n", err, msg)
			}
		}
	}()

	<-exitSignalChannel
	log.Println("Gracefully stopping...")
	producer.Close()
	consumer.Close()
	log.Println("...stopped")
}