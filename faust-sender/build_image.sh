#!/bin/bash
cp ../common/models/event.py ./event.py
docker build -t faust_sender --network=host .
