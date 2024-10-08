package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "0.0.0.0:50051"

func main() {
	cc, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}

	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(cc)

	client := pb.NewCalculatorServiceClient(cc)
	//doSum(client, 2, 3)
	//getPrimes(client, 120)
	//getAvg(client, []int32{1, 13, 33})
	//doMax(client, []int{4, -5, 10, 0, -3, 34})
	doSqrt(client, -3)
}
