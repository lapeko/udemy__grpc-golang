package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"io"
	"log"
)

func GetPrimes(client pb.PrimesStreamingServiceClient, num uint32) {
	stream, err := client.GetPrimes(context.Background(), &pb.PrimeRequest{Number: num})

	if err != nil {
		log.Fatalln(err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln(err)
		}

		log.Println(res)
	}
}
