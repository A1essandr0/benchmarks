#!/bin/bash
cp ../common/models/event.go ./event.go
docker build -t franz_collector --network=host .
