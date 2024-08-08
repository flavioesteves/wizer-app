package main

import (
	"fmt"
	"net"

	pb "github.com/flavioesteves/wizer-app/user/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var USER_SERVICE_HOST = "0.0.0.0:50051"

type Server struct {
	pb.UserServiceServer
}

func main() {
	fmt.Println("User Service started")

	listen, err := net.Listen("tcp", USER_SERVICE_HOST)

	fmt.Printf("Listen: %v\n", listen.Addr().String())

	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
