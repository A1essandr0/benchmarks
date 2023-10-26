package main

import "github.com/lovoo/goka"

var (
	brokers = []string{"kafka:9092"}
	port = "0.0.0.0:5005"
	topic goka.Stream = "test-events-from"
)