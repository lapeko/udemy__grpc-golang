package main

import (
	"context"
	"errors"
	"fmt"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
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

func (s *greetServer) GreetEveryone(stream grpc.BidiStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalln(err)
		}

		name := req.Name
		log.Println("Request received... Name:", name)

		if err := stream.Send(&pb.GreetResponse{Response: fmt.Sprintf("Hello, %s", name)}); err != nil {
			log.Fatalln(err)
		}
	}
}

func (s *greetServer) GreetDeadline(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	for i := 0; i < 3; i++ {

		log.Println("Request proceeding...")

		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			log.Println("Client canceled request")
			return nil, status.Error(codes.DeadlineExceeded, "The client canceled the request")
		}
		time.Sleep(1 * time.Second)
	}

	log.Println("Sending response...")

	return &pb.GreetResponse{Response: fmt.Sprintf("Hello, %s", req.Name)}, nil
}
