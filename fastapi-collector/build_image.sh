#!/bin/bash
cp ../common/models/event.py ./event.py
cp ../common/libs/kafka_client.py ./kafka_client.py
docker build -t fastapi_collector --network=host .
