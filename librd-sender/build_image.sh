#!/bin/bash
cp ../common/models/event.go ./event.go
docker build -t librd_sender --network=host .
