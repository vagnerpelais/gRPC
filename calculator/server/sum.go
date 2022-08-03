package main

import (
	"context"
	"log"

	pb "nappsolutions.io/google-admin-v2/calculator/proto"
)

func (s *Servers) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked %v\n", in)

	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
