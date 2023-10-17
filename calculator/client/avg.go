package main

import (
	"context"
	"log"
	"time"

	pb "github.com/mgnsharon/grpc-go-course/calculator/proto"
)

func getAvg(c pb.CalculatorServiceClient) {
	log.Println("getAvg invoked")

	reqs := []*pb.AvgRequest{
		{ Num: 1 },
		{ Num: 2 },
		{ Num: 3 },
		{ Num: 4 },
	}

	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("Error while calling getAvg %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response from getAvg: %v\n", err)
	}

	log.Printf("Avg: %v\n", res.Result)
}