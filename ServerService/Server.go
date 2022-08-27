package main

import (
	pb "ServerService/gRPCproto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPasswordsStorageServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetStorage(ctx context.Context, in *pb.Token) (*pb.Storage, error) {
	log.Printf("Received: %X", in.GetToken())
	defer ctx.Done()
	return &pb.Storage{
		Site:     []string{"MySite", "1"},
		Login:    []string{"Hello"},
		Password: []string{"00000"},
	}, nil
}

func main() {
	//TODO: why using flag? want to have cmd access?
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9999))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPasswordsStorageServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
