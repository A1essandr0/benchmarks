package main


type Config struct {
	brokers []string
	sourceTopic string
	destinationTopic string
	consumerGroupName string
}

var config = Config{
	brokers: []string{"kafka:9092"},
	sourceTopic: "test-events-from",
	destinationTopic: "test-events-to",
	consumerGroupName: "franz-sender",
}