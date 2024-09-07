package main

import (
	"context"
	"fmt"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func (s *greetServer) Greet(ctx context.Context, greet *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Response: fmt.Sprintf("Hello %s", greet.Name),
	}, nil
}

func (s *greetServer) GreetManyTimes(gr *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.GreetResponse{Response: fmt.Sprintf("Hello %s. Iteration %d", gr.Name, i+1)}); err != nil {
			return err
		}
	}
	return nil
}

func (s *greetServer) GreetLong(clientStream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	responseString := ""

	for {
		req, err := clientStream.Recv()

		if err == io.EOF {
			if err := clientStream.SendAndClose(&pb.GreetResponse{Response: responseString}); err != nil {
				log.Fatalln(err)
			}
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Request received.", req)

		responseString += fmt.Sprintf("Hello, %s\n", req.Name)
	}
}
