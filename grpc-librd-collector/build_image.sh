#!/bin/bash
cp -r ../common/proto ./
docker build -t grpc_librd_collector --network=host .
