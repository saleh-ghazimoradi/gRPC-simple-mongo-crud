package main

import (
	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookItem struct {
	ID       primitive.ObjectID `bson:"_id"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBook(data *BookItem) *pb.Book {
	return &pb.Book{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
