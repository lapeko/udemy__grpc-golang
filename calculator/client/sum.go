package main

import (
	"context"
	"github.com/lapeko/udemy__grpc-golang/sum/proto"
	pb "github.com/lapeko/udemy__grpc-golang/sum/proto"
	"log"
)

func doSum(ssc proto.SumServiceClient, num1 int, num2 int) {
	log.Println("Sending request...")

	res, err := ssc.Sum(context.Background(), &pb.SumRequest{Num1: int32(num1), Num2: int32(num2)})

	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Response received:\nThe sum of %d and %d is: %d", num1, num2, res.Sum)
}
