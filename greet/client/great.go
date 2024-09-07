package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"time"
)

func doGreat(c pb.GreetServiceClient) {
	log.Println("Client sending request...")

	greet, err := c.Greet(context.Background(), &pb.GreetRequest{
		Name: "Vitali",
	})

	if err != nil {
		log.Println("Error", err)
	}

	log.Println("Client received response: ", greet.Response)
}

func doGreetList(c pb.GreetServiceClient) {
	log.Println("Client sending request...")

	stream, err := c.GreetManyTimes(context.Background(), &pb.GreetRequest{Name: "Vital"})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		response, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(response)
	}

	log.Println("Client response received...")
}

func doGreetLong(c pb.GreetServiceClient) {
	names := []string{"Maria", "Anna", "Karina"}

	stream, err := c.GreetLong(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, name := range names {
		if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(1 * time.Second)
	}

	response, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(response.Response)
}

func doGreetEveryone(s pb.GreetServiceClient, names []string) {
	stream, err := s.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, name := range names {
			time.Sleep(1 * time.Second)

			log.Println("Request sending... Name:", name)

			if err := stream.Send(&pb.GreetRequest{Name: name}); err != nil {
				log.Fatalln(err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				break
			}

			if err != nil {
				log.Fatalln(err)
			}

			log.Printf("Response received: %s\n", res.Response)
		}
	}()

	<-waitc
}

func doDeadlineGreet(s pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := s.GreetDeadline(ctx, &pb.GreetRequest{Name: "Revolution"})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Fatalln("request timeout exceeded")
			}
			log.Fatalln("unexpected gRPC error", err)
		}
		log.Fatalln("Not gRPC error", err)
	}

	log.Printf("Response successfully received: %s\n", res.Response)
}
