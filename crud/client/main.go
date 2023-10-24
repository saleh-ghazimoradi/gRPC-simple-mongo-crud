package main

import (
	"fmt"
	"log"

	pb "github.com/saleh-ghazimoradi/gRPC-simple-mongo-crud/crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Couldn't connect to client: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBookShareServiceClient(conn)

	fmt.Println(c)
}
