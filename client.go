package main

import (
	"context"
	"log"

	"github.com/shehzabdeveloper/gRPC-rest/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dail(":9000", grpc.withInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.close()

	c := echo.NewEchoServiceClient(conn)
	response, err := c.Echo(context.Background(), &echo.EchoMessage{Value: "Check client"})
	if err != nil {
		log.Fatalf("Error when calling Echo:: %s", err)
	}
	log.Printf("Response from server: %s", response.Value)
}
