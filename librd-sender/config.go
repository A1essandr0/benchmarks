package main


type Config struct {
	broker string
	sourceTopic string
	destinationTopic string
	consumerGroupName string
}

var config = Config{
	broker: "kafka:9092",
	sourceTopic: "test-events-from",
	destinationTopic: "test-events-to",
	consumerGroupName: "librd-sender",
}