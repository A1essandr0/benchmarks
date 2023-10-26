#!/bin/bash
cp ../common/models/event.go ./event.go
docker build -t franz_sender --network=host .
