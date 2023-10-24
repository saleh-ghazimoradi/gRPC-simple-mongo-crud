package main

import (
	"context"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
)

func readBook(c pb.BookShareServiceClient, id string) *pb.Book {
	log.Println("----> readBook was executed <----")

	req := &pb.BookId{
		Id: id,
	}

	res, err := c.ReadBook(context.Background(), req)

	if err != nil {
		log.Printf("error happened while reading: %v\n", err)

	}

	log.Printf("Book was read: %v\n", res)
	return res
}
