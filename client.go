package main

import (
	"context"
	"fmt"
	pb "go-zero-demo/tuser"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "dev-user-rpc:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()
	print(456)
	c := pb.NewUserClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	print(123)
	defer cancel()
	r, err := c.Login(ctx, &pb.LoginRequest{Email: "andy@indochat.co.id", Password: "123456"})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}
	fmt.Println("Response:", r.Email)
}
