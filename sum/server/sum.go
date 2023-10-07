package main

import (
	"context"
	"log"

	pb "github.com/mgnsharon/grpc-go-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.NumOne + in.NumTwo,
	}, nil
}