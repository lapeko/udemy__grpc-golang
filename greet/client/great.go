package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
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
