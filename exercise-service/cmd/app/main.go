package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/flavioesteves/wizer-app/exercise/cmd/app/handler"
	"github.com/flavioesteves/wizer-app/exercise/internal/database"
	pb "github.com/flavioesteves/wizer-app/exercise/proto"
)

var EXERCISE_SERVICE_HOST = "0.0.0.0:50054"

func main() {
	fmt.Println("Exercise Service started")

	dbConn, err := database.ConnectToDB()

	if err != nil {
		log.Panic("Can't connect to PostGres!")
	}

	// Init Server
	serverConfig := handler.NewServerConfig(dbConn)
	listen, err := net.Listen("tcp", EXERCISE_SERVICE_HOST)
	fmt.Printf("Listen %v\n", listen.Addr().String())

	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterExerciseServiceServer(s, serverConfig)
	reflection.Register(s)

	if err = s.Serve(listen); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
	}

}
