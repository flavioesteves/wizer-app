package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/flavioesteves/wizer-app/authentication/cmd/app/handler"
	"github.com/flavioesteves/wizer-app/authentication/internal/database"
	"github.com/flavioesteves/wizer-app/authentication/internal/middleware"
	pb "github.com/flavioesteves/wizer-app/authentication/proto"
)

var AUTH_SERVICE_HOST = "0.0.0.0:50055"

func main() {
	fmt.Println("Authentication Service Started")

	// connect to Redis
	redisClient, err := middleware.RedisConn()
	if err != nil {
		fmt.Printf("Error to connect to redis %v\n", err)
	}
	fmt.Println(redisClient)

	// connect to Database
	dbConn, err := database.ConnectToDB()

	if err != nil {
		log.Panic("Can't connect to PostGres!")
	}

	// Init Server
	serverConfig := handler.NewServerConfig(redisClient, dbConn)
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
