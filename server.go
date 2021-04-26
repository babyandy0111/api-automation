package main

import (
	"context"
	pb "go-zero-demo/tuser"
	"google.golang.org/grpc"
	"log"
	"net"
)

type service struct {
	pb.UnimplementedUserServer
}

func (s *service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Response, error) {
	log.Printf("Received: %v", in.Email)
	return &pb.Response{Email: "Hello"}, nil
}

func main() {
	addr := "127.0.0.1:9999"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Server listening on", addr)
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServer(gRPCServer, &service{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
