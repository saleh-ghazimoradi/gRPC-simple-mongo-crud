package main

import (
	"context"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
)

func createBook(c pb.BookShareServiceClient) string {
	log.Println("----> createBook <----")

	book := &pb.Book{
		AuthorId: "Saleh Ghazimoradi",
		Title:    "A Pamphlet By Which Success Would Be In The Cards!",
		Content:  "This is a pamphlet containing some informative information and techniques which help you ace the entrance exam",
	}

	res, err := c.CreateBook(context.Background(), book)

	if err != nil {
		log.Fatalf("unexpected error creating book %v\n:", err)
	}

	log.Printf("Blog has been created successfully. Id: %s\n", res.Id)

	return res.Id
}
