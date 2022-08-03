package main

import (
	"context"
	"log"

	pb "nappsolutions.io/google-admin-v2/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
