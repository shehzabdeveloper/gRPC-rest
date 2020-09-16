package main

import (
	"fmt"
	"log"
	"net"

	"github.com/shehzabdeveloper/gRPC-rest/echo"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Server started")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	ss := echo.Server{}
	grpcServer := grpc.NewServer()

	echo.RegisterEchoServiceService(grpcServer, &ss)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
