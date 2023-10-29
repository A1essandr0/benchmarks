package main


type Config struct {
	port string
	broker string
	sourceTopic string
}

var config = Config{
	port: "0.0.0.0:5005",
	broker: "kafka:9092",
	sourceTopic: "test-events-from",
}
