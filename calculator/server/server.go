package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const addr = "0.0.0.0:50051"

type calculatorServer struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &calculatorServer{})

	log.Println("Server is running on ", addr)
	log.Fatalln(s.Serve(lis))
}
