package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"io"
	"log"
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

func doGreetList(c pb.GreetListServiceClient) {
	log.Println("Client sending request...")

	stream, err := c.GreetList(context.Background(), &pb.GreetRequest{Name: "Vital"})

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
