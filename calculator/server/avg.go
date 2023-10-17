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
	var res float64 = 0
	var item_count = float64(len(nums))
	for _, num := range nums {
		res += float64(num)
	}
	if item_count > 0 {
		res /= item_count
	} else {
		log.Fatal("Error: No items found.")
	}
	return res
}
