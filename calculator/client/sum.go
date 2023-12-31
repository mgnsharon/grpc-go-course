package main

import (
	"context"
	"log"

	pb "github.com/mgnsharon/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		NumOne: 23,
		NumTwo: 32,
	})

	if err != nil {
		log.Fatalf("Could not calculate the sum: %v\n", err)
	}

	log.Printf("Sum: %v", res.Result)
}