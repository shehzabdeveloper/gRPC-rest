package echo

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, in *EchoMessage) (*EchoMessage, error) {
	log.Printf("Receive message body from client: %s", in.Value)
	return &EchoMessage{Value: "Hello From the Server!"}, nil
}
