package main

import (
	"context"

	pb "github.com/akifkadioglu/protocol-buffers-golang/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Text: "Hello " + in.Id,
	}, nil
}
