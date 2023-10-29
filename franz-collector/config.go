package main


type Config struct {
	port string
	brokers []string
	sourceTopic string
	consumerGroupName string
}

var config = Config{
	port: "0.0.0.0:5005",
	brokers: []string{"kafka:9092"},
	sourceTopic: "test-events-from",
	consumerGroupName: "franz-collector",
}
