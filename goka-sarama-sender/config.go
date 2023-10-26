package main

import "github.com/lovoo/goka"

var (
	groupName goka.Group = "goka-sarama-sender"
	brokers = []string{"kafka:9092"}
	sourceTopic goka.Stream = "test-events-from"
	processedTopic goka.Stream = "test-events-to"	
)