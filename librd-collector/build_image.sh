#!/bin/bash
cp ../common/models/event.go ./event.go
docker build -t librd_collector --network=host .
