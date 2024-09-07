package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
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

func (s *calculatorServer) Max(stream grpc.BidiStreamingServer[pb.MaxRequest, pb.MaxResponse]) error {
	initialised := false
	var maxNumber int

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		reqNumber := int(req.Number)

		if !initialised {
			initialised = true
			maxNumber = reqNumber
		}

		if reqNumber > maxNumber {
			maxNumber = reqNumber
		}

		if err := stream.Send(&pb.MaxResponse{MaxNumber: int32(maxNumber)}); err != nil {
			log.Fatalln(err)
		}
	}
}

func (s *calculatorServer) Sqrt(c context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	if in.Number <= 0 {
		return nil, status.Error(codes.InvalidArgument, "PLease, provide positive value")
	}

	num := math.Sqrt(float64(in.Number))

	return &pb.SqrtResponse{Number: num}, nil
}
