package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"log"
)

func (s *sumServer) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Request received with params: %v\n", in)

	return &pb.SumResponse{
		Sum: in.Num1 + in.Num2,
	}, nil
}
