package main

import (
	"context"
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func doMax(c pb.CalculatorServiceClient, numbers []int) {
	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalln(err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, number := range numbers {
			time.Sleep(500 * time.Millisecond)
			log.Println("Sending request to server. Number:", number)

			if err := stream.Send(&pb.MaxRequest{Number: int32(number)}); err != nil {
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

			log.Println("Response received. Max:", res.MaxNumber)
		}
	}()

	<-waitc
}

func doSqrt(c pb.CalculatorServiceClient, num int32) {
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: num})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.InvalidArgument {
				log.Fatalln("You probably provided not positive value")
			}
			log.Fatalf("Error. Status code %d. %s", e.Code(), e.Message())
		}
		log.Fatalln(err)
	}

	log.Printf("Sqrt of %d is %.2f\n", num, res.Number)
}
