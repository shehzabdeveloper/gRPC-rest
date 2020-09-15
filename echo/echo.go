package echo

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, in *EchoMessage) (*EchoMessage, error) {
	log.Printf("Client message: %s", in.Value)
	return &EchoMessage{Value: "Hello from Server"}, nil
}
