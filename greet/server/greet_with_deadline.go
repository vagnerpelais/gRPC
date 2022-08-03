package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/internal/status"
	"google.golang.org/grpc/status"
	pb "nappsolutions.io/google-admin-v2/greet/proto"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request")
			return nil, status.Error(codes.Canceled, "The client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
