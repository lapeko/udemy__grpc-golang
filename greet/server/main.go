package main

import (
	pb "github.com/lapeko/udemy__grpc-golang/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	address = "0.0.0.0:50051"
)

type server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	log.Println("Server running on ", address)
	log.Fatalln(s.Serve(lis))
}
