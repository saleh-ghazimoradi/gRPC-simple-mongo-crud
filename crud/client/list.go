package main

import (
	"context"
	"io"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func listBooks(c pb.BookShareServiceClient) {
	log.Println("----> listBooks <----")

	stream, err := c.ListBooks(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("error while calling listBooks: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("something happened: %v", err)
		}

		log.Println(res)
	}
}
