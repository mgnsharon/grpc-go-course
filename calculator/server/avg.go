package main

import (
	"io"
	"log"

	pb "github.com/mgnsharon/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	nums := []int64{}

	for {
		req, err := stream.Recv()

		if err == io.EOF {
		
			return stream.SendAndClose(&pb.AvgResponse{
				Result: calculateAvg(nums),
			})
		}

		if err != nil {
			log.Fatalf("Error calculating avg: %v\n", err)
		}
		nums = append(nums, req.Num)
		
	}

}

func calculateAvg(nums []int64) float64 {
	var sum float64 = 0
	var item_count = float64(len(nums))

	if item_count < 1 {
		log.Fatal("Error: No items found.")
	}
	for _, num := range nums {
		sum += float64(num)
	}
	
	return sum/item_count
}
