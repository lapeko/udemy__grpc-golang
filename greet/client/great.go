package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
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
