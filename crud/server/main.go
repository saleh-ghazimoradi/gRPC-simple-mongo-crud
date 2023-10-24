package main

import (
	"log"
	"net"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BookShareServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBookShareServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
