package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"log"
)

func (s *primesServer) GetPrimes(request *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
	num := request.Number
	prime := uint32(2)

	for prime <= num {
		if num%prime == 0 {
			num /= prime
			err := stream.Send(&pb.PrimeResponse{
				Prime: prime,
			})
			if err != nil {
				log.Fatalln(err)
			}
			continue
		}
		prime++
	}

	return nil
}
