package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/flavioesteves/wizer-app/routine/cmd/app/handler"
	"github.com/flavioesteves/wizer-app/routine/internal/database"
	pb "github.com/flavioesteves/wizer-app/routine/proto"
)

var ROUTINE_SERVICE_HOST = "0.0.0.0:50053"

func main() {
	fmt.Println("User Service started")

	dbConn, err := database.ConnectToDB()

	if err != nil {
		log.Panic("Can't connect to PostGres!")
	}

	// Init Server
	serverConfig := handler.NewServerConfig(dbConn)
	listen, err := net.Listen("tcp", ROUTINE_SERVICE_HOST)
	fmt.Printf("Listen %v\n", listen.Addr().String())

	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRoutineServiceServer(s, serverConfig)
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}
}
