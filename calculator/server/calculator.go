package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func (s *calculatorServer) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Request received with params: %v\n", in)

	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}

func (s *calculatorServer) Primes(request *pb.PrimeRequest, stream grpc.ServerStreamingServer[pb.PrimeResponse]) error {
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

func (s *calculatorServer) Avg(client grpc.ClientStreamingServer[pb.AvgRequest, pb.AvgResponse]) error {
	log.Println("Avg function call...")

	numbers := make([]int32, 0)

	for {
		req, err := client.Recv()

		if err == io.EOF {
			var sum float32
			for _, num := range numbers {
				sum += float32(num)
			}
			avg := sum / float32(len(numbers))
			if err := client.SendAndClose(&pb.AvgResponse{Avg: avg}); err != nil {
				log.Fatalln(err)
			}
			log.Printf("Avg function response: %f\n", avg)
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Printf("New value received: %d\n", req.Number)

		numbers = append(numbers, req.Number)
	}
}
