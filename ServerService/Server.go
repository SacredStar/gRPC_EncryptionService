package main

import (
	"context"
	"flag"
	"fmt"
	pb "gRPCClientServerForEncryption/proto3"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPasswordsStorageServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) GetStorage(ctx context.Context, in *pb.Token) (*pb.Storage, error) {
	log.Printf("Received: %v", in.GetToken())
	return &pb.Storage{
		Site:     []string{"MySite", "1"},
		Login:    []string{"Hello"},
		Password: []string{"00000"},
	}, nil
}

func main() {
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
