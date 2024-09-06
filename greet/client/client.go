package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const address = "0.0.0.0:50051"

func main() {
	con, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}

	defer con.Close()
}
