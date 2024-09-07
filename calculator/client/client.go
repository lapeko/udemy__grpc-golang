package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/sum/proto"
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

	sumClient := pb.NewSumServiceClient(cc)
	doSum(sumClient, 2, 3)
}
