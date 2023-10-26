package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	pb "grpc-librd-collector/proto"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
)

var (
	p *kafka.Producer
)


type Server struct {
	pb.UnimplementedCollectorServiceServer
}

func (s *Server) PostRawEvent(ctx context.Context, in *pb.RawEvent) (*pb.Response, error) {
	log.Printf("received: %+v", in)
	serialized, err := json.Marshal(in)
	// serialized, err := proto.Marshal(in)
	if err != nil {		
		log.Println("error while serializing", err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic_from, Partition: kafka.PartitionAny},
		Value:          serialized,
	}, nil)
	if err != nil {
		log.Println("couldn't send to kafka")
	}
	log.Println("...sent to kafka")

	return &pb.Response{Status: "OK"}, nil
}


func main() {
	var err error
	p, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		panic(err)
	}
	
	log.Printf("producer: %+v, topic: %s\n", p, topic_from)
	log.Printf("starting gRPC server on tcp %s\n", tcp_port)
	listener, err := net.Listen("tcp", tcp_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterCollectorServiceServer(server, &Server{})
	// reflection.Register(server)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

