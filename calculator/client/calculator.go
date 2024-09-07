package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"io"
	"log"
	"time"
)

func doSum(client pb.CalculatorServiceClient, num1 int, num2 int) {
	log.Println("Sending request...")

	res, err := client.Sum(context.Background(), &pb.SumRequest{Num1: int32(num1), Num2: int32(num2)})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Response received:\nThe sum of %d and %d is: %d", num1, num2, res.Sum)
}

func getPrimes(client pb.CalculatorServiceClient, num uint32) {
	stream, err := client.Primes(context.Background(), &pb.PrimeRequest{Number: num})

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

func getAvg(client pb.CalculatorServiceClient, numbers []int32) {
	stream, err := client.Avg(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	for _, number := range numbers {
		if err := stream.Send(&pb.AvgRequest{Number: number}); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Response received: Avg is: %f\n", res.Avg)
}
