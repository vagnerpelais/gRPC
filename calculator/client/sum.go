package main

import (
	"context"
	"log"

	pb "nappsolutions.io/google-admin-v2/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 1,
	})

	if err != nil {
		log.Fatalf("Could no sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
