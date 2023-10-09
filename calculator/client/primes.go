package main

import (
	"context"
	"io"
	"log"

	pb "github.com/mgnsharon/grpc-go-course/calculator/proto"
)

func getPrimes(c pb.CalculatorServiceClient) {
	log.Println("getPrimes was invoked")

	req := &pb.PrimesRequest{
		Num: 120,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while getting Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}
		log.Printf("Prime: %v\n", res.Result)
	}
}