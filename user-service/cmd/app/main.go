package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/flavioesteves/wizer-app/user/cmd/app/handler"
	"github.com/flavioesteves/wizer-app/user/internal/database"
	pb "github.com/flavioesteves/wizer-app/user/proto"
)

var USER_SERVICE_HOST = "0.0.0.0:50052"

func main() {
	fmt.Println("User Service started")
	// Connect to database
	dbConn, err := database.ConnectToDB()

	if err != nil {
		log.Panic("Can't connect to PostGres!")
	}

	// Init Server
	serverConfig := handler.NewServerConfig(dbConn)
	listen, err := net.Listen("tcp", USER_SERVICE_HOST)
	fmt.Printf("Listen: %v\n", listen.Addr().String())

	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, serverConfig)
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
	}
}
