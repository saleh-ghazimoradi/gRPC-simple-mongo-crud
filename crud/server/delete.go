package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBook(ctx context.Context, in *pb.BookId) (*emptypb.Empty, error) {
	log.Printf("deleteBook was invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"cannot parse id",
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("cannot delete object in mongo: %v\n", err),
		)
	}

	if res.DeletedCount == 0 {

		return nil, status.Errorf(
			codes.NotFound,
			"book was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
