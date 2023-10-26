#!/bin/bash
grpcurl -proto grpc-librd-collector/proto/server.proto list
grpcurl -proto grpc-librd-collector/proto/server.proto describe
grpcurl -proto grpc-librd-collector/proto/server.proto -d '{"source": "id_1234", "event_name": "starting event" }' -plaintext localhost:5006 go_grpc_collector_proto.CollectorService/PostRawEvent