package main

import (
	"log"

	pb "nappsolutions.io/google-admin-v2/calculator/proto"
)

func (s *Servers) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with: %v\n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}
