package main

import (
	"context"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
)

func deleteBook(c pb.BookShareServiceClient, id string) {
	log.Println("----> deleteBook <----")

	_, err := c.DeleteBook(context.Background(), &pb.BookId{Id: id})

	if err != nil {
		log.Fatalf("error while deleting book: %v", err)
	}

	log.Println("book was successfully deleted")
}
