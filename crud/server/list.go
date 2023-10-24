package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBooks(in *emptypb.Empty, stream pb.BookShareService_ListBooksServer) error {
	log.Println("listBooks was executed")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BookItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while decoding data from mongo: %v", err),
			)
		}

		stream.Send(documentToBook(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknown internal error: %v", err),
		)
	}

	return nil
}
