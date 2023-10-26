#!/bin/bash
cp ../common/models/event.go ./event.go
docker build -t goka_sarama_collector --network=host .
