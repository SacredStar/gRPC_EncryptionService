package main

import (
	"context"
	"fmt"
	pb "gRPCClientServerForEncryption/proto3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "google.golang.org/protobuf/proto"
	"log"
	"time"
)

var (
	addr = fmt.Sprintf("%s", "localhost:9999")
)

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPasswordsStorageClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetStorage(ctx, &pb.Token{
		Token: "123",
	})
	if err != nil {
		log.Fatalf("could not Get Strage: %v", err)
	}
	log.Printf("Greeting: %s", r.GetLogin())

}
