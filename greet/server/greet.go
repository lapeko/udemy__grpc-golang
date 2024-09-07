package main

import (
	"context"
	"fmt"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
)

func (s *greetServer) Greet(ctx context.Context, greet *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Response: fmt.Sprintf("Hello %s", greet.Name),
	}, nil
}

func (s *greetListServer) GreetList(gr *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.GreetResponse{Response: fmt.Sprintf("Hello %s. Iteration %d", gr.Name, i+1)}); err != nil {
			return err
		}
	}
	return nil
}
