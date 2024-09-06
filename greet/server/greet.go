package main

import (
	"context"
	"fmt"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Response: fmt.Sprintf("Hello %s", in.Name),
	}, nil
}
