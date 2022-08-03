package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "nappsolutions.io/google-admin-v2/greet/proto"
)

func doGreetWithDeadlines(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "Bohm",
	}

	res, err := c.GreetWithDeadline(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			} else {
				log.Fatalf("Unexpected gRPC error: %v\n", e)
			}
		} else {
			log.Fatalf("A non gRPC error: %s\n", err)
		}
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
