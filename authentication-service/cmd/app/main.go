package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/flavioesteves/wizer-app/authentication/cmd/app/handler"
	"github.com/flavioesteves/wizer-app/authentication/internal/middleware"
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
)

var AUTH_SERVICE_HOST = "0.0.0.0:50055"

func main() {
	fmt.Println("Authentication Service Started")

	// connect to Redis
	redis, err := middleware.RedisConn()
	if err != nil {
		fmt.Printf("Error to connect to redis %v\n", err)
	}

	fmt.Println(redis)

	// Init Server
	serverConfig := handler.NewServerConfig()
	listen, err := net.Listen("tcp", AUTH_SERVICE_HOST)
	fmt.Printf("Listen: %v\n", listen.Addr().String())

	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(s, serverConfig)
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
