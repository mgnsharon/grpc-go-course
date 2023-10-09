package main

import (
	"fmt"
	"log"

	pb "github.com/mgnsharon/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(req *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", req)
	var k int64 = 2

	getFactor(req.Num, k, stream)

	return nil
}

func getFactor(n, k int64, stream pb.CalculatorService_PrimesServer) {
	
	if n % k == 0 {
		fmt.Println(k)
		stream.Send(&pb.PrimesResponse{
			Result: k,
		})
		n /= k
		if n <= 1 {
			fmt.Println("Done")
		} else {
			getFactor(n, k, stream)
		}
	} else {
		getFactor(n, k+1, stream)
	}

}