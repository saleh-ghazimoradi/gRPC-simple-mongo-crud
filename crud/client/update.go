package main

import (
	"context"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
)

func updateBook(c pb.BookShareServiceClient, id string) {
	log.Println("----> updateBook was executed <----")

	newBook := &pb.Book{
		Id:       id,
		AuthorId: "A new author",
		Title:    "New version of the book",
		Content:  "some amazing stuff has been added to the previous version of the book",
	}

	_, err := c.UpdateBook(context.Background(), newBook)

	if err != nil {
		log.Fatalf("error happened while updating: %v\n", err)
	}

	log.Println("boos was updated successfully!")
}
